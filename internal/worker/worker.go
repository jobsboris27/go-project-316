package worker

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code/internal/checker"
	"code/internal/parser"
	"code/internal/report"
	"code/internal/shared"
)

type Config struct {
	UserAgent   string
	Timeout     time.Duration
	Retries     int
	Concurrency int
	Depth       int
	HTTPClient  *http.Client
}

type Result struct {
	Page         report.PageReport
	Links        []string
	ShouldQueue  bool
	NextDepth    int
}

func ProcessURL(ctx context.Context, pageURL string, depth int, rootHost string, cfg Config, checkerCfg checker.Config) Result {
	normalizedURL := shared.NormalizeURLFromItem(pageURL)
	
	page := report.PageReport{
		URL:          normalizedURL,
		Depth:        depth,
		DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		page.Status = "error"
		page.Error = err.Error()
		page.SEO = &report.SEOReport{}
		page.BrokenLinks = nil
		page.Assets = nil
		return Result{Page: page}
	}

	if cfg.UserAgent != "" {
		req.Header.Set("User-Agent", cfg.UserAgent)
	}

	var resp *http.Response
	var lastErr error
	maxAttempts := cfg.Retries + 1

	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			select {
			case <-time.After(100 * time.Millisecond):
			case <-ctx.Done():
				page.Status = "error"
				page.Error = ctx.Err().Error()
				page.SEO = &report.SEOReport{}
				page.BrokenLinks = nil
				page.Assets = nil
				return Result{Page: page}
			}
		}

		resp, err = cfg.HTTPClient.Do(req)
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
		page.Status = "error"
		if lastErr != nil {
			page.Error = lastErr.Error()
		} else {
			page.Error = err.Error()
		}
		page.HTTPStatus = 0
		page.SEO = &report.SEOReport{}
		page.BrokenLinks = nil
		page.Assets = nil
		return Result{Page: page}
	}
	defer func() { _ = resp.Body.Close() }()

	page.HTTPStatus = resp.StatusCode
	page.Status = "ok"

	contentType := resp.Header.Get("Content-Type")

	if strings.Contains(contentType, "text/html") {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			page.SEO = (*report.SEOReport)(parser.ParseSEO(bytes.NewReader(body)))
			pageURLParsed, _ := url.Parse(pageURL)

			links := parser.ExtractLinks(string(body), pageURLParsed)
			brokenLinks := checker.CheckLinks(ctx, pageURL, body, checkerCfg)
			for _, bl := range brokenLinks {
				page.BrokenLinks = append(page.BrokenLinks, report.BrokenLink(bl))
			}
			assets := checker.CheckAssets(ctx, pageURL, body, checkerCfg)
			for _, a := range assets {
				page.Assets = append(page.Assets, report.Asset(a))
			}

			page.BrokenLinks = dedupeBrokenLinks(page.BrokenLinks)
			page.Assets = dedupeAssets(page.Assets)

			return Result{
				Page:        page,
				Links:       links,
				ShouldQueue: depth+1 < cfg.Depth,
				NextDepth:   depth + 1,
			}
		}
	} else if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			page.SEO = (*report.SEOReport)(parser.ParseSEO(bytes.NewReader(body)))
		}
	}

	if page.BrokenLinks == nil {
		page.BrokenLinks = []report.BrokenLink{}
	}
	if page.Assets == nil {
		page.Assets = []report.Asset{}
	}
	if page.SEO == nil {
		page.SEO = &report.SEOReport{}
	}

	return Result{Page: page}
}

func dedupeBrokenLinks(links []report.BrokenLink) []report.BrokenLink {
	seen := make(map[string]bool)
	result := make([]report.BrokenLink, 0)
	for _, link := range links {
		if !seen[link.URL] {
			seen[link.URL] = true
			result = append(result, link)
		}
	}
	return result
}

func dedupeAssets(assets []report.Asset) []report.Asset {
	seen := make(map[string]bool)
	result := make([]report.Asset, 0)
	for _, asset := range assets {
		if !seen[asset.URL] {
			seen[asset.URL] = true
			result = append(result, asset)
		}
	}
	return result
}
