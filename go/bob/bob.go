// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Rule out any empty expressions
	if len(remark) <= 0 || !strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890?") {
		return "Fine. Be that way!"
	}

	// Trim whitespace
	remark = strings.TrimSpace(remark)

	if strings.HasSuffix(remark, "?") {
		if strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ToUpper(remark) == remark {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else if strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ToUpper(remark) == remark {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}
