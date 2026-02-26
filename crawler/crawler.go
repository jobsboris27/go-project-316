package crawler

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"code/internal/checker"
	"code/internal/parser"
	"code/internal/report"
	"code/internal/shared"
)

type job struct {
	url   string
	depth int
}

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
	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{
			Timeout: opts.Timeout,
		}
	}

	if opts.Concurrency <= 0 {
		opts.Concurrency = 4
	}

	rootURL, err := url.Parse(opts.URL)
	if err != nil {
		return report.Marshal(report.NewErrorReport(opts.URL, opts.Depth, err.Error()), opts.IndentJSON)
	}
	rootHost := rootURL.Host

	rep := report.Report{
		RootURL:     opts.URL,
		Depth:       opts.Depth,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Pages:       make([]report.PageReport, 0),
	}

	if err := ctx.Err(); err != nil {
		return report.Marshal(rep, opts.IndentJSON)
	}

	visited := make(map[string]bool)
	var visitedMu sync.Mutex

	visited[shared.NormalizeURL(rootURL)] = true

	limiter := shared.NewRateLimiter(opts.Delay, opts.RPS)
	if limiter != nil {
		defer limiter.Stop()
	}

	jobChan := make(chan job, 100)
	resultChan := make(chan report.PageReport, 100)

	var wg sync.WaitGroup
	for i := 0; i < opts.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, opts, rootHost, limiter, jobChan, resultChan)
		}()
	}

	jobChan <- job{url: opts.URL, depth: 0}
	pending := 1
	done := false

	for !done && pending > 0 {
		select {
		case res := <-resultChan:
			rep.Pages = append(rep.Pages, res)

			if res.Status == "ok" && res.Depth < opts.Depth && res.RawBody != nil && len(res.RawBody) > 0 {
				links, _ := parser.ParseHTMLLinks(bytes.NewReader(res.RawBody))
				baseURL, _ := url.Parse(res.URL)

				for _, link := range links {
					absoluteURL := shared.ResolveURL(baseURL, link)
					if absoluteURL == "" || !shared.IsValidScheme(absoluteURL) {
						continue
					}

					linkURL, err := url.Parse(absoluteURL)
					if err != nil {
						continue
					}

					if linkURL.Host != rootHost {
						continue
					}

					visitedMu.Lock()
					if !visited[shared.NormalizeURL(linkURL)] {
						newDepth := res.Depth + 1
						if newDepth > opts.Depth {
							visitedMu.Unlock()
							continue
						}
						visited[shared.NormalizeURL(linkURL)] = true
						select {
						case jobChan <- job{url: absoluteURL, depth: newDepth}:
							pending++
						case <-ctx.Done():
							visitedMu.Unlock()
							done = true
							break
						}
					}
					visitedMu.Unlock()

					if done {
						break
					}
				}
			}
			if !done {
				pending--
			}

		case <-ctx.Done():
			done = true
		}
	}

	close(jobChan)
	wg.Wait()
	close(resultChan)

	report.Sort(rep)

	seen := make(map[string]bool)
	uniquePages := make([]report.PageReport, 0, len(rep.Pages))
	for _, page := range rep.Pages {
		if !seen[page.URL] {
			seen[page.URL] = true
			uniquePages = append(uniquePages, page)
		}
	}
	rep.Pages = uniquePages

	return report.Marshal(rep, opts.IndentJSON)
}

func worker(ctx context.Context, opts Options, rootHost string, limiter *shared.RateLimiter, jobChan <-chan job, resultChan chan<- report.PageReport) {
	for j := range jobChan {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if j.depth > opts.Depth {
			continue
		}

		if limiter != nil {
			if err := limiter.Wait(ctx); err != nil {
				return
			}
		}

		pageReport := analyzePage(ctx, opts, j.url, j.depth)

		select {
		case resultChan <- pageReport:
		case <-ctx.Done():
			return
		}
	}
}

func analyzePage(ctx context.Context, opts Options, pageURL string, depth int) report.PageReport {
	pageReport := report.NewPageReport(pageURL, depth)

	if err := ctx.Err(); err != nil {
		pageReport.Status = "error"
		pageReport.Error = err.Error()
		pageReport.SEO = &SEOReport{}
		return pageReport
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		pageReport.Status = "error"
		pageReport.Error = err.Error()
		pageReport.SEO = &SEOReport{}
		return pageReport
	}

	if opts.UserAgent != "" {
		req.Header.Set("User-Agent", opts.UserAgent)
	}

	var resp *http.Response
	var lastErr error
	maxAttempts := opts.Retries + 1

	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			select {
			case <-time.After(100 * time.Millisecond):
			case <-ctx.Done():
				pageReport.Status = "error"
				pageReport.Error = ctx.Err().Error()
				return pageReport
			}
		}

		resp, err = opts.HTTPClient.Do(req)
		if err == nil {
			if resp.StatusCode < 500 && resp.StatusCode != 429 {
				break
			}
			if resp.StatusCode >= 500 || resp.StatusCode == 429 {
				lastErr = nil
				continue
			}
			break
		}

		lastErr = err
	}

	if err != nil {
		pageReport.Status = "error"
		if lastErr != nil {
			pageReport.Error = lastErr.Error()
		} else {
			pageReport.Error = err.Error()
		}
		pageReport.SEO = &SEOReport{}
		return pageReport
	}
	defer func() { _ = resp.Body.Close() }()

	pageReport.HTTPStatus = resp.StatusCode
	pageReport.Status = "ok"

	contentType := resp.Header.Get("Content-Type")

	checkerCfg := checker.Config{
		UserAgent:   opts.UserAgent,
		Timeout:     opts.Timeout,
		Concurrency: opts.Concurrency,
		HTTPClient:  opts.HTTPClient,
	}

	if pageReport.Status == "ok" {
		if strings.Contains(contentType, "text/html") {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				pageReport.RawBody = body
				pageReport.SEO = (*report.SEOReport)(parser.ParseSEO(bytes.NewReader(body)))
				pageURLParsed, _ := url.Parse(pageURL)
				assetInfos := parser.ExtractAssets(string(body), pageURLParsed)
				for _, ai := range assetInfos {
					pageReport.Assets = append(pageReport.Assets, Asset{
						URL:  ai.URL,
						Type: ai.Type,
					})
				}
				pageReport.BrokenLinks = toBrokenLinks(checker.CheckLinks(ctx, pageURL, body, checkerCfg))
				pageReport.Assets = toAssets(checker.CheckAssets(ctx, pageURL, body, checkerCfg))
			}
		} else if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				pageReport.RawBody = body
				pageReport.SEO = (*report.SEOReport)(parser.ParseSEO(bytes.NewReader(body)))
			}
		}

		if pageReport.BrokenLinks == nil {
			pageReport.BrokenLinks = make([]BrokenLink, 0)
		}
		if pageReport.Assets == nil {
			pageReport.Assets = make([]Asset, 0)
		}
		if pageReport.SEO == nil {
			pageReport.SEO = &SEOReport{}
		}
	}

	return pageReport
}

func toBrokenLinks(links []checker.BrokenLink) []BrokenLink {
	result := make([]BrokenLink, len(links))
	for i, link := range links {
		result[i] = BrokenLink{
			URL:        link.URL,
			StatusCode: link.StatusCode,
			Error:      link.Error,
		}
	}
	return result
}

func toAssets(assets []checker.Asset) []Asset {
	result := make([]Asset, len(assets))
	for i, asset := range assets {
		result[i] = Asset{
			URL:        asset.URL,
			Type:       asset.Type,
			StatusCode: asset.StatusCode,
			SizeBytes:  asset.SizeBytes,
			Error:      asset.Error,
		}
	}
	return result
}
