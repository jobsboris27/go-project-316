package crawler

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"code/internal/shared"
)

type job struct {
	url   string
	depth int
}

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
		return shared.MarshalReport(createErrorReport(opts.URL, opts.Depth, err.Error()), opts.IndentJSON)
	}
	rootHost := rootURL.Host

	report := Report{
		RootURL:     opts.URL,
		Depth:       opts.Depth,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Pages:       make([]PageReport, 0),
	}

	if err := ctx.Err(); err != nil {
		return shared.MarshalReport(report, opts.IndentJSON)
	}

	visited := make(map[string]bool)
	var visitedMu sync.Mutex
	
	visited[shared.NormalizeURL(rootURL)] = true

	limiter := shared.NewRateLimiter(opts.Delay, opts.RPS)
	if limiter != nil {
		defer limiter.Stop()
	}

	jobChan := make(chan job, 100)
	resultChan := make(chan PageReport, 100)

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
			report.Pages = append(report.Pages, res)

			if res.Status == "ok" && res.Depth < opts.Depth && res.RawBody != nil && len(res.RawBody) > 0 {
				parser := NewParser(opts.HTTPClient, opts.UserAgent, opts.Timeout)
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
						visited[shared.NormalizeURL(linkURL)] = true
						select {
						case jobChan <- job{url: absoluteURL, depth: res.Depth + 1}:
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

	sortReport(report)

	return shared.MarshalReport(report, opts.IndentJSON)
}

func sortReport(report Report) {
	sort.Slice(report.Pages, func(i, j int) bool {
		if report.Pages[i].Depth != report.Pages[j].Depth {
			return report.Pages[i].Depth < report.Pages[j].Depth
		}
		return report.Pages[i].URL < report.Pages[j].URL
	})

	for i := range report.Pages {
		sortPage(&report.Pages[i])
	}
}

func sortPage(page *PageReport) {
	sort.Slice(page.Assets, func(i, j int) bool {
		if page.Assets[i].Type != page.Assets[j].Type {
			return page.Assets[i].Type < page.Assets[j].Type
		}
		return page.Assets[i].URL < page.Assets[j].URL
	})

	sort.Slice(page.BrokenLinks, func(i, j int) bool {
		return page.BrokenLinks[i].URL < page.BrokenLinks[j].URL
	})
}

func worker(ctx context.Context, opts Options, rootHost string, limiter *shared.RateLimiter, jobChan <-chan job, resultChan chan<- PageReport) {
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

		pageReport := analyzePage(ctx, opts, j.url, j.depth, rootHost)

		select {
		case resultChan <- pageReport:
		case <-ctx.Done():
			return
		}
	}
}

func createErrorReport(rootURL string, depth int, errMsg string) Report {
	return Report{
		RootURL:     rootURL,
		Depth:       depth,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Pages: []PageReport{
			{
				URL:          rootURL,
				Depth:        0,
				Status:       "error",
				Error:        errMsg,
				DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
				SEO:          &SEOReport{},
			},
		},
	}
}

func analyzePage(ctx context.Context, opts Options, pageURL string, depth int, rootHost string) PageReport {
	pageReport := PageReport{
		URL:          pageURL,
		Depth:        depth,
		DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
	}

	if err := ctx.Err(); err != nil {
		pageReport.Status = "error"
		pageReport.Error = err.Error()
		return pageReport
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		pageReport.Status = "error"
		pageReport.Error = err.Error()
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
		return pageReport
	}
	defer func() { _ = resp.Body.Close() }()

	pageReport.HTTPStatus = resp.StatusCode
	pageReport.Status = "ok"

	contentType := resp.Header.Get("Content-Type")
	parser := NewParser(opts.HTTPClient, opts.UserAgent, opts.Timeout)

	if strings.Contains(contentType, "text/html") {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			pageReport.RawBody = body
			pageReport.SEO = parser.ParseSEO(bytes.NewReader(body))
			pageReport.BrokenLinks = parser.CheckLinks(ctx, pageURL, body, rootHost, opts.Concurrency)
			pageReport.Assets = parser.CheckAssets(ctx, pageURL, body, opts.Concurrency)
		}
	} else if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			pageReport.RawBody = body
			pageReport.SEO = parser.ParseSEO(bytes.NewReader(body))
			pageReport.BrokenLinks = make([]BrokenLink, 0)
			pageReport.Assets = make([]Asset, 0)
		}
	} else {
		pageReport.RawBody = nil
	}

	if pageReport.Status == "ok" {
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
	// Для error страниц оставляем nil чтобы в JSON было null

	return pageReport
}
