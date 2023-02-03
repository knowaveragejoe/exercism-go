package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, c := range word {
		// Skip hypens and spaces
		if string(c) == "-" || string(c) == " " {
			continue
		}

		if strings.Count(word, string(c)) >= 2 {
			return false
		}
	}

	return true
}
