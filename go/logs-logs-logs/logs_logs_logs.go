package logs

import (
	"fmt"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	// var application string
	for _, c := range log {
		char := fmt.Sprintf("%U", c)
		if char == "U+2757" {
			return "recommendation"
		} else if char == "U+1F50D" {
			return "search"
		} else if char == "U+2600" {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
