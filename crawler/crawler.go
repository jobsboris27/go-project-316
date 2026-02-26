package crawler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"code/internal/shared"

	"golang.org/x/net/html"
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
		return marshalReport(createErrorReport(opts.URL, opts.Depth, err.Error()), opts.IndentJSON)
	}
	rootHost := rootURL.Host

	report := Report{
		RootURL:     opts.URL,
		Depth:       opts.Depth,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Pages:       make([]PageReport, 0),
	}

	if err := ctx.Err(); err != nil {
		return marshalReport(report, opts.IndentJSON)
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

	for {
		select {
		case res := <-resultChan:
			report.Pages = append(report.Pages, res)

			if res.Status == "ok" && res.Depth < opts.Depth && res.RawBody != nil {
				links, _ := parseHTMLLinks(bytes.NewReader(res.RawBody))
				baseURL, _ := url.Parse(res.URL)

				for _, link := range links {
					absoluteURL := resolveURL(baseURL, link)
					if absoluteURL == "" || !isValidScheme(absoluteURL) {
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
							goto drain
						}
					}
					visitedMu.Unlock()
				}
			}
			pending--

		case <-ctx.Done():
			goto drain
		}

		if pending == 0 {
			break
		}
	}

drain:
	close(jobChan)
	wg.Wait()
	close(resultChan)

	return marshalReport(report, opts.IndentJSON)
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
				URL:         rootURL,
				Depth:       0,
				Status:      "error",
				Error:       errMsg,
				BrokenLinks: make([]BrokenLink, 0),
				SEO:         &SEOReport{},
				Assets:      make([]Asset, 0),
			},
		},
	}
}

func analyzePage(ctx context.Context, opts Options, pageURL string, depth int, rootHost string) PageReport {
	pageReport := PageReport{
		URL:          pageURL,
		Depth:        depth,
		DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
		BrokenLinks:  make([]BrokenLink, 0),
		SEO:          &SEOReport{},
		Assets:       make([]Asset, 0),
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
	if strings.Contains(contentType, "text/html") {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			pageReport.RawBody = body
			pageReport.SEO = parseSEO(bytes.NewReader(body))
			links := checkLinks(ctx, opts, pageURL, body, rootHost)
			pageReport.BrokenLinks = append(pageReport.BrokenLinks, links...)
			assets := checkAssets(ctx, opts, pageURL, body)
			pageReport.Assets = append(pageReport.Assets, assets...)
		}
	}

	return pageReport
}

func parseHTMLLinks(r io.Reader) ([]string, error) {
	var links []string

	tokenizer := html.NewTokenizer(r)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			}
			return links, tokenizer.Err()
		}

		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" && attr.Val != "" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}

	return links, nil
}

func resolveURL(base *url.URL, link string) string {
	if link == "" {
		return ""
	}

	parsed, err := url.Parse(link)
	if err != nil {
		return ""
	}

	if parsed.IsAbs() {
		return link
	}

	resolved := base.ResolveReference(parsed)
	return resolved.String()
}

func isValidScheme(u string) bool {
	parsed, err := url.Parse(u)
	if err != nil {
		return false
	}

	if parsed.Scheme == "" {
		return false
	}

	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func checkLinks(ctx context.Context, opts Options, pageURL string, body []byte, rootHost string) []BrokenLink {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return nil
	}

	links, _ := parseHTMLLinks(bytes.NewReader(body))
	absoluteLinks := make([]string, 0)

	for _, link := range links {
		absoluteURL := resolveURL(baseURL, link)
		if absoluteURL != "" && isValidScheme(absoluteURL) {
			absoluteLinks = append(absoluteLinks, absoluteURL)
		}
	}

	if len(absoluteLinks) == 0 {
		return make([]BrokenLink, 0)
	}

	semaphore := make(chan struct{}, opts.Concurrency)
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

			brokenLink := checkLink(ctx, opts, linkURL)
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

func checkLink(ctx context.Context, opts Options, linkURL string) *BrokenLink {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, linkURL, nil)
	if err != nil {
		return &BrokenLink{
			URL:   linkURL,
			Error: shared.FormatError("failed to create request: %v", err),
		}
	}

	if opts.UserAgent != "" {
		req.Header.Set("User-Agent", opts.UserAgent)
	}

	client := opts.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: opts.Timeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		getReq, getErr := http.NewRequestWithContext(ctx, http.MethodGet, linkURL, nil)
		if getErr == nil {
			if opts.UserAgent != "" {
				getReq.Header.Set("User-Agent", opts.UserAgent)
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

func parseSEO(r io.Reader) *SEOReport {
	seo := &SEOReport{}

	var titleText, descriptionText string
	var foundTitle, foundDescription, foundH1 bool

	tokenizer := html.NewTokenizer(r)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			break
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()

			switch token.Data {
			case "title":
				if tokenizer.Next() == html.TextToken {
					titleText = strings.TrimSpace(tokenizer.Token().Data)
					foundTitle = true
				}

			case "meta":
				var name, content string
				for _, attr := range token.Attr {
					if attr.Key == "name" {
						name = strings.ToLower(attr.Val)
					}
					if attr.Key == "content" {
						content = attr.Val
					}
				}
				if name == "description" {
					descriptionText = strings.TrimSpace(content)
					foundDescription = true
				}

			case "h1":
				foundH1 = true
			}
		}
	}

	seo.HasTitle = foundTitle
	seo.Title = titleText
	seo.HasDescription = foundDescription
	seo.Description = descriptionText
	seo.HasH1 = foundH1

	return seo
}

func marshalReport(report Report, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(report, "", "  ")
	}
	return json.Marshal(report)
}

type assetTag struct {
	URL  string
	Type string
}

func parseAssets(r io.Reader) []assetTag {
	var assets []assetTag

	tokenizer := html.NewTokenizer(r)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			break
		}

		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			token := tokenizer.Token()

			switch token.Data {
			case "img":
				for _, attr := range token.Attr {
					if attr.Key == "src" && attr.Val != "" {
						assets = append(assets, assetTag{URL: attr.Val, Type: "image"})
					}
				}

			case "script":
				for _, attr := range token.Attr {
					if attr.Key == "src" && attr.Val != "" {
						assets = append(assets, assetTag{URL: attr.Val, Type: "script"})
					}
				}

			case "link":
				var rel, href string
				for _, attr := range token.Attr {
					if attr.Key == "rel" {
						rel = strings.ToLower(attr.Val)
					}
					if attr.Key == "href" {
						href = attr.Val
					}
				}
				if rel == "stylesheet" && href != "" {
					assets = append(assets, assetTag{URL: href, Type: "style"})
				}
			}
		}
	}

	return assets
}

func checkAssets(ctx context.Context, opts Options, pageURL string, body []byte) []Asset {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return make([]Asset, 0)
	}

	assetTags := parseAssets(bytes.NewReader(body))
	if len(assetTags) == 0 {
		return make([]Asset, 0)
	}

	assetCache := make(map[string]Asset)
	assets := make([]Asset, 0)
	var mu sync.Mutex

	semaphore := make(chan struct{}, opts.Concurrency)
	var wg sync.WaitGroup

	for _, tag := range assetTags {
		absoluteURL := resolveURL(baseURL, tag.URL)
		if absoluteURL == "" || !isValidScheme(absoluteURL) {
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

			asset := checkAsset(ctx, opts, assetURL, assetType)

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

func checkAsset(ctx context.Context, opts Options, assetURL, assetType string) Asset {
	asset := Asset{
		URL:  assetURL,
		Type: assetType,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, assetURL, nil)
	if err != nil {
		asset.Error = err.Error()
		return asset
	}

	if opts.UserAgent != "" {
		req.Header.Set("User-Agent", opts.UserAgent)
	}

	client := opts.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: opts.Timeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		getReq, getErr := http.NewRequestWithContext(ctx, http.MethodGet, assetURL, nil)
		if getErr == nil {
			if opts.UserAgent != "" {
				getReq.Header.Set("User-Agent", opts.UserAgent)
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
