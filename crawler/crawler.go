package crawler

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	"code/internal/checker"
	"code/internal/report"
	"code/internal/shared"
	"code/internal/worker"
)

type Options struct {
	URL         string
	Depth       int
	Retries     int
	Delay       time.Duration
	Timeout     time.Duration
	UserAgent   string
	Concurrency int
	IndentJSON  bool
	RPS         int
	HTTPClient  *http.Client
}

type BrokenLink = report.BrokenLink
type Asset = report.Asset
type SEOReport = report.SEOReport
type PageReport = report.PageReport
type Report = report.Report

func Analyze(ctx context.Context, opts Options) ([]byte, error) {
	cfg := normalizeOptions(&opts)

	rootURL, err := shared.ParseAndValidateURL(opts.URL)
	if err != nil {
		return report.Marshal(report.NewErrorReport(opts.URL, opts.Depth, err.Error()), opts.IndentJSON)
	}

	normalizedRootURL := shared.NormalizeURL(rootURL)

	visited := shared.NewVisitedSet()
	queue := shared.NewURLQueue(normalizedRootURL)
	semaphore := make(chan struct{}, opts.Concurrency)
	var wg sync.WaitGroup

	builder := report.NewBuilder(rootURL.String(), opts.Depth)

	checkerCfg := checker.Config{
		UserAgent:   opts.UserAgent,
		Timeout:     opts.Timeout,
		Concurrency: opts.Concurrency,
		HTTPClient:  opts.HTTPClient,
	}

	for {
		item := queue.Dequeue()

		if item == nil {
			wg.Wait()
			if queue.IsEmpty() {
				break
			}
			continue
		}

		normalizedItemURL := shared.NormalizeURLFromItem(item.URL)

		if visited.Contains(normalizedItemURL) {
			continue
		}
		visited.Add(normalizedItemURL)

		wg.Add(1)
		semaphore <- struct{}{}

		go func(jobURL string, jobDepth int) {
			defer wg.Done()
			defer func() { <-semaphore }()

			normalizedJobURL := shared.NormalizeURLFromItem(jobURL)
			result := worker.ProcessURL(ctx, normalizedJobURL, jobDepth, rootURL.Host, cfg, checkerCfg)
			builder.AddPage(result.Page)

			if result.ShouldQueue {
				toAdd := make([]shared.URLWithDepth, 0)
				for _, link := range result.Links {
					linkURL, err := url.Parse(link)
					if err != nil || linkURL.Host != rootURL.Host {
						continue
					}
					normalized := shared.NormalizeURL(linkURL)
					if !visited.Contains(normalized) {
						toAdd = append(toAdd, shared.URLWithDepth{URL: link, Depth: result.NextDepth})
					}
				}
				if len(toAdd) > 0 {
					queue.Enqueue(toAdd)
				}
			}
		}(item.URL, item.Depth)
	}

	return builder.Encode(opts.IndentJSON)
}

func normalizeOptions(opts *Options) worker.Config {
	if opts.Concurrency <= 0 {
		opts.Concurrency = 4
	}
	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{
			Timeout: opts.Timeout,
		}
	}

	return worker.Config{
		UserAgent:   opts.UserAgent,
		Timeout:     opts.Timeout,
		Retries:     opts.Retries,
		Concurrency: opts.Concurrency,
		Depth:       opts.Depth,
		HTTPClient:  opts.HTTPClient,
	}
}
