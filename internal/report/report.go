package report

import (
	"encoding/json"
	"sort"
	"sync"
	"time"
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
	Error      string `json:"error,omitempty"`
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
	BrokenLinks   []BrokenLink `json:"broken_links"`
	DiscoveredAt  string       `json:"discovered_at"`
	SEO           *SEOReport   `json:"seo"`
	Assets        []Asset      `json:"assets"`
}

type Report struct {
	RootURL     string       `json:"root_url"`
	Depth       int          `json:"depth"`
	GeneratedAt string       `json:"generated_at"`
	Pages       []PageReport `json:"pages"`
}

type Builder struct {
	report *Report
	mu     sync.Mutex
}

func NewBuilder(rootURL string, depth int) *Builder {
	return &Builder{
		report: &Report{
			RootURL:     rootURL,
			Depth:       depth,
			GeneratedAt: time.Now().UTC().Format(time.RFC3339),
			Pages:       []PageReport{},
		},
	}
}

func (b *Builder) AddPage(page PageReport) {
	if page.Error != "" {
		page.BrokenLinks = nil
		page.Assets = nil
	} else {
		if page.BrokenLinks == nil {
			page.BrokenLinks = []BrokenLink{}
		}
		if page.Assets == nil {
			page.Assets = []Asset{}
		}
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	b.report.Pages = append(b.report.Pages, page)
}

func (b *Builder) Encode(indent bool) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	sort.SliceStable(b.report.Pages, func(i, j int) bool {
		if b.report.Pages[i].Depth != b.report.Pages[j].Depth {
			return b.report.Pages[i].Depth < b.report.Pages[j].Depth
		}
		return b.report.Pages[i].URL < b.report.Pages[j].URL
	})

	for i := range b.report.Pages {
		sort.SliceStable(b.report.Pages[i].Assets, func(j, k int) bool {
			return b.report.Pages[i].Assets[j].Type < b.report.Pages[i].Assets[k].Type
		})
		sort.SliceStable(b.report.Pages[i].BrokenLinks, func(j, k int) bool {
			return b.report.Pages[i].BrokenLinks[j].URL < b.report.Pages[i].BrokenLinks[k].URL
		})
	}

	if indent {
		return json.MarshalIndent(b.report, "", "  ")
	}
	return json.Marshal(b.report)
}

func NewPageReport(url string, depth int) PageReport {
	return PageReport{
		URL:          url,
		Depth:        depth,
		DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
	}
}

func NewErrorReport(rootURL string, depth int, errMsg string) Report {
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
				BrokenLinks:  nil,
				DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
				SEO:          &SEOReport{},
				Assets:       nil,
			},
		},
	}
}

func Marshal(report Report, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(report, "", "  ")
	}
	return json.Marshal(report)
}

func Sort(rep Report) {
	sort.Slice(rep.Pages, func(i, j int) bool {
		if rep.Pages[i].Depth != rep.Pages[j].Depth {
			return rep.Pages[i].Depth < rep.Pages[j].Depth
		}
		return rep.Pages[i].URL < rep.Pages[j].URL
	})

	for i := range rep.Pages {
		sort.Slice(rep.Pages[i].Assets, func(j, k int) bool {
			if rep.Pages[i].Assets[j].Type != rep.Pages[i].Assets[k].Type {
				return rep.Pages[i].Assets[j].Type < rep.Pages[i].Assets[k].Type
			}
			return rep.Pages[i].Assets[j].URL < rep.Pages[i].Assets[k].URL
		})
		sort.Slice(rep.Pages[i].BrokenLinks, func(j, k int) bool {
			return rep.Pages[i].BrokenLinks[j].URL < rep.Pages[i].BrokenLinks[k].URL
		})
	}
}
