package pangram

import (
	"strings"
	"testing"
)

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "empty sentence",
		input:       "",
		expected:    false,
	},
	{
		description: "perfect lower case",
		input:       "abcdefghijklmnopqrstuvwxyz",
		expected:    true,
	},
	{
		description: "only lower case",
		input:       "the quick brown fox jumps over the lazy dog",
		expected:    true,
	},
	{
		description: "missing the letter 'x'",
		input:       "a quick movement of the enemy will jeopardize five gunboats",
		expected:    false,
	},
	{
		description: "missing the letter 'h'",
		input:       "five boxing wizards jump quickly at it",
		expected:    false,
	},
	{
		description: "with underscores",
		input:       "the_quick_brown_fox_jumps_over_the_lazy_dog",
		expected:    true,
	},
	{
		description: "with numbers",
		input:       "the 1 quick brown fox jumps over the 2 lazy dogs",
		expected:    true,
	},
	{
		description: "missing letters replaced by numbers",
		input:       "7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog",
		expected:    false,
	},
	{
		description: "mixed case and punctuation",
		input:       "\"Five quacking Zephyrs jolt my wax bed.\"",
		expected:    true,
	},
	{
		description: "a-m and A-M are 26 different characters but not a pangram",
		input:       "abcdefghijklm ABCDEFGHIJKLM",
		expected:    false,
	},
}

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

func BenchmarkPangram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsPangram(test.input)
		}
	}
}
