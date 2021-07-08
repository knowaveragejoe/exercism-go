// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate returns an abbreviation of a string
func Abbreviate(s string) string {
	var abbr string

	// Replace hyphens with spaces, underscores with nothing
	noHyphens := strings.ReplaceAll(s, "-", " ")
	noUnderscores := strings.ReplaceAll(noHyphens, "_", "")
	words := strings.Split(noUnderscores, " ") // Split on spaces


	for _, word := range words {
		if len(word) <= 0 {
			continue
		}
		abbr += strings.ToUpper(string(word[0]))
	}

	return abbr
}

0