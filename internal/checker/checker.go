package checker

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"code/internal/parser"
	"code/internal/shared"
)

type BrokenLink struct {
	URL        string `json:"url"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type Asset struct {
	URL        string `json:"url"`
	Type       string `json:"type"`
	StatusCode int    `json:"status_code"`
	SizeBytes  int64  `json:"size_bytes"`
	Error      string `json:"error"`
}

type Config struct {
	UserAgent   string
	Timeout     time.Duration
	Concurrency int
	HTTPClient  *http.Client
}

func CheckLinks(ctx context.Context, pageURL string, body []byte, cfg Config) []BrokenLink {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return nil
	}

	links, _ := parser.ParseHTMLLinks(bytes.NewReader(body))
	absoluteLinks := make([]string, 0)

	for _, link := range links {
		absoluteURL := shared.ResolveURL(baseURL, link)
		if absoluteURL != "" && shared.IsValidScheme(absoluteURL) {
			absoluteLinks = append(absoluteLinks, absoluteURL)
		}
	}

	if len(absoluteLinks) == 0 {
		return make([]BrokenLink, 0)
	}

	semaphore := make(chan struct{}, cfg.Concurrency)
	var wg sync.WaitGroup
	var brokenMu sync.Mutex
	broken := make([]BrokenLink, 0)

	for _, link := range absoluteLinks {
		select {
		case <-ctx.Done():
			continue
		default:
		}

		semaphore <- struct{}{}
		wg.Add(1)

		go func(linkURL string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			brokenLink := CheckLink(ctx, linkURL, cfg)
			if brokenLink != nil {
				brokenMu.Lock()
				broken = append(broken, *brokenLink)
				brokenMu.Unlock()
			}
		}(link)
	}

	wg.Wait()

	return broken
}

func CheckLink(ctx context.Context, linkURL string, cfg Config) *BrokenLink {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, linkURL, nil)
	if err != nil {
		return &BrokenLink{
			URL:   linkURL,
			Error: shared.FormatError("failed to create request: %v", err),
		}
	}

	if cfg.UserAgent != "" {
		req.Header.Set("User-Agent", cfg.UserAgent)
	}

	client := cfg.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: cfg.Timeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		getReq, getErr := http.NewRequestWithContext(ctx, http.MethodGet, linkURL, nil)
		if getErr == nil {
			if cfg.UserAgent != "" {
				getReq.Header.Set("User-Agent", cfg.UserAgent)
			}
			getResp, getErr := client.Do(getReq)
			if getErr != nil {
				return &BrokenLink{
					URL:   linkURL,
					Error: err.Error(),
				}
			}
			defer func() { _ = getResp.Body.Close() }()

			if getResp.StatusCode >= 400 {
				return &BrokenLink{
					URL:        linkURL,
					StatusCode: getResp.StatusCode,
				}
			}
			return nil
		}

		return &BrokenLink{
			URL:   linkURL,
			Error: err.Error(),
		}
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 400 {
		return &BrokenLink{
			URL:        linkURL,
			StatusCode: resp.StatusCode,
		}
	}

	return nil
}

func CheckAssets(ctx context.Context, pageURL string, body []byte, cfg Config) []Asset {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return make([]Asset, 0)
	}

	assetTags := parser.ParseAssets(bytes.NewReader(body))
	if len(assetTags) == 0 {
		return make([]Asset, 0)
	}

	assetCache := make(map[string]Asset)
	assets := make([]Asset, 0)
	var mu sync.Mutex

	semaphore := make(chan struct{}, cfg.Concurrency)
	var wg sync.WaitGroup

	for _, tag := range assetTags {
		absoluteURL := shared.ResolveURL(baseURL, tag.URL)
		if absoluteURL == "" || !shared.IsValidScheme(absoluteURL) {
			continue
		}

		mu.Lock()
		if _, exists := assetCache[absoluteURL]; exists {
			mu.Unlock()
			continue
		}
		mu.Unlock()

		semaphore <- struct{}{}
		wg.Add(1)

		go func(assetURL, assetType string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			asset := CheckAsset(ctx, assetURL, assetType, cfg)

			mu.Lock()
			if _, exists := assetCache[assetURL]; !exists {
				assetCache[assetURL] = asset
				assets = append(assets, asset)
			}
			mu.Unlock()
		}(absoluteURL, tag.Type)
	}

	wg.Wait()

	return assets
}

func CheckAsset(ctx context.Context, assetURL, assetType string, cfg Config) Asset {
	asset := Asset{
		URL:  assetURL,
		Type: assetType,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, assetURL, nil)
	if err != nil {
		asset.Error = err.Error()
		return asset
	}

	if cfg.UserAgent != "" {
		req.Header.Set("User-Agent", cfg.UserAgent)
	}

	client := cfg.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: cfg.Timeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		getReq, getErr := http.NewRequestWithContext(ctx, http.MethodGet, assetURL, nil)
		if getErr == nil {
			if cfg.UserAgent != "" {
				getReq.Header.Set("User-Agent", cfg.UserAgent)
			}
			getResp, getErr := client.Do(getReq)
			if getErr == nil {
				defer func() { _ = getResp.Body.Close() }()
				body, _ := io.ReadAll(getResp.Body)
				asset.StatusCode = getResp.StatusCode
				asset.SizeBytes = int64(len(body))
				if getResp.StatusCode >= 400 {
					asset.Error = http.StatusText(getResp.StatusCode)
				}
				return asset
			}
		}
		asset.Error = err.Error()
		return asset
	}
	defer func() { _ = resp.Body.Close() }()

	asset.StatusCode = resp.StatusCode

	contentLength := resp.Header.Get("Content-Length")
	if contentLength != "" {
		if size, err := strconv.ParseInt(contentLength, 10, 64); err == nil {
			asset.SizeBytes = size
		}
	} else {
		body, _ := io.ReadAll(resp.Body)
		asset.SizeBytes = int64(len(body))
	}

	if resp.StatusCode >= 400 {
		asset.Error = http.StatusText(resp.StatusCode)
	}

	return asset
}
