// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var result []string
	for i, w := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
			break
		}
		result = append(result, fmt.Sprintf("For want of a %v the %v was lost.", w, rhyme[i+1]))
	}

	return result
}
