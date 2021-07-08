package scrabble

import "strings"

func Score(word string) int {
	word = strings.ToUpper(word)
	scrabbleScores := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	var score int

	// Iterate through letters in the word.
	for _, letter := range word {
		// Test if the letters match are in the keys of scrabble scores.
		for l, s := range scrabbleScores {
			if strings.Contains(l, string(letter)) {
				score += s
				break
			}
		}
	}

	return score
}
