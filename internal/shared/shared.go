package shared

import (
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
