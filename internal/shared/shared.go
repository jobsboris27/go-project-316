package shared

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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

func MarshalReport(report interface{}, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(report, "", "  ")
	}
	return json.Marshal(report)
}
