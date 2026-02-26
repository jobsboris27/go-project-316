package crawler

import (
	"net/http"
	"time"
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

type SEOReport struct {
	HasTitle       bool   `json:"has_title"`
	Title          string `json:"title"`
	HasDescription bool   `json:"has_description"`
	Description    string `json:"description"`
	HasH1          bool   `json:"has_h1"`
}

type PageReport struct {
	URL           string       `json:"url"`
	Depth         int          `json:"depth"`
	HTTPStatus    int          `json:"http_status"`
	Status        string       `json:"status"`
	Error         string       `json:"error,omitempty"`
	BrokenLinks   []BrokenLink `json:"broken_links,omitempty"`
	DiscoveredAt  string       `json:"discovered_at,omitempty"`
	SEO           *SEOReport   `json:"seo,omitempty"`
	Assets        []Asset      `json:"assets,omitempty"`
	RawBody       []byte       `json:"-"`
}

type Report struct {
	RootURL     string       `json:"root_url"`
	Depth       int          `json:"depth"`
	GeneratedAt string       `json:"generated_at"`
	Pages       []PageReport `json:"pages"`
}
