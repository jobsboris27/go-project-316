package crawler

import (
	"bytes"
	"context"
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

type assetTag struct {
	URL  string
	Type string
}

type Parser struct {
	httpClient *http.Client
	userAgent  string
	timeout    time.Duration
}

func NewParser(client *http.Client, userAgent string, timeout time.Duration) *Parser {
	return &Parser{
		httpClient: client,
		userAgent:  userAgent,
		timeout:    timeout,
	}
}

func (p *Parser) ParseHTMLLinks(r io.Reader) ([]string, error) {
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

func (p *Parser) ParseAssets(r io.Reader) []assetTag {
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

func (p *Parser) ParseSEO(r io.Reader) *SEOReport {
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

func (p *Parser) CheckLinks(ctx context.Context, pageURL string, body []byte, rootHost string, concurrency int) []BrokenLink {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return nil
	}

	links, _ := p.ParseHTMLLinks(bytes.NewReader(body))
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

	semaphore := make(chan struct{}, concurrency)
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

			brokenLink := p.CheckLink(ctx, linkURL)
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

func (p *Parser) CheckLink(ctx context.Context, linkURL string) *BrokenLink {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, linkURL, nil)
	if err != nil {
		return &BrokenLink{
			URL:   linkURL,
			Error: shared.FormatError("failed to create request: %v", err),
		}
	}

	if p.userAgent != "" {
		req.Header.Set("User-Agent", p.userAgent)
	}

	client := p.httpClient
	if client == nil {
		client = &http.Client{Timeout: p.timeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		getReq, getErr := http.NewRequestWithContext(ctx, http.MethodGet, linkURL, nil)
		if getErr == nil {
			if p.userAgent != "" {
				getReq.Header.Set("User-Agent", p.userAgent)
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

func (p *Parser) CheckAssets(ctx context.Context, pageURL string, body []byte, concurrency int) []Asset {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return make([]Asset, 0)
	}

	assetTags := p.ParseAssets(bytes.NewReader(body))
	if len(assetTags) == 0 {
		return make([]Asset, 0)
	}

	assetCache := make(map[string]Asset)
	assets := make([]Asset, 0)
	var mu sync.Mutex

	semaphore := make(chan struct{}, concurrency)
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

			asset := p.CheckAsset(ctx, assetURL, assetType)

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

func (p *Parser) CheckAsset(ctx context.Context, assetURL, assetType string) Asset {
	asset := Asset{
		URL:  assetURL,
		Type: assetType,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, assetURL, nil)
	if err != nil {
		asset.Error = err.Error()
		return asset
	}

	if p.userAgent != "" {
		req.Header.Set("User-Agent", p.userAgent)
	}

	client := p.httpClient
	if client == nil {
		client = &http.Client{Timeout: p.timeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		getReq, getErr := http.NewRequestWithContext(ctx, http.MethodGet, assetURL, nil)
		if getErr == nil {
			if p.userAgent != "" {
				getReq.Header.Set("User-Agent", p.userAgent)
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
