package shared

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
)

type URLWithDepth struct {
	URL   string
	Depth int
}

type URLQueue struct {
	items []URLWithDepth
	mu    sync.Mutex
}

func NewURLQueue(rootURL string) *URLQueue {
	return &URLQueue{
		items: []URLWithDepth{{URL: rootURL, Depth: 0}},
	}
}

func (q *URLQueue) Enqueue(urls []URLWithDepth) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, urls...)
}

func (q *URLQueue) Dequeue() *URLWithDepth {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return &item
}

func (q *URLQueue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}

type VisitedSet struct {
	urls map[string]bool
	mu   sync.Mutex
}

func NewVisitedSet() *VisitedSet {
	return &VisitedSet{
		urls: make(map[string]bool),
	}
}

func (v *VisitedSet) Add(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.urls[url] = true
}

func (v *VisitedSet) Contains(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.urls[url]
}

func FormatError(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NormalizeURL(u *url.URL) string {
	normalized := u.String()
	if u.Path == "" || u.Path == "/" {
		normalized = u.Scheme + "://" + u.Host + "/"
	}
	return normalized
}

func NormalizeURLFromItem(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	return NormalizeURL(u)
}

func ResolveURL(base *url.URL, link string) string {
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

func IsValidScheme(u string) bool {
	parsed, err := url.Parse(u)
	if err != nil {
		return false
	}

	if parsed.Scheme == "" {
		return false
	}

	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func ParseAndValidateURL(urlStr string) (*url.URL, error) {
	return url.Parse(urlStr)
}

func MarshalReport(report interface{}, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(report, "", "  ")
	}
	return json.Marshal(report)
}
