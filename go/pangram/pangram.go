package pangram

import "strings"

func IsPangram(input string) bool {
	allLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	letters := make(map[string]int)

	// Count letters
	for _, l := range input {
		letters[strings.ToUpper(string(l))] += 1
	}

	// Loop over all letters and check they exist at least once
	for _, l := range allLetters {
		_, ok := letters[l]
		if !ok {
			return false
		}
	}

	return true
}
