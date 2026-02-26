package shared

import "fmt"

func FormatError(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
