package luhn

import (
	"strings"
	"regexp"
	"strconv"
)

// Determines if a given number(as a string) passes the Luhn validation algorithm
func Valid(numString string) bool {
	// Strip spaces from the string
	numString = strings.ReplaceAll(numString, " ", "")

	// Valid numbers must be greater than 1 digit
	if len(numString) <= 1 {
		return false
	}

	// If there are any non-digit characters after this, the string is invalid
	nonDigits := regexp.MustCompile(`\D+`)
	if nonDigits.MatchString(numString) {
		return false
	}

	var sum int

	for j, i := 0, len(numString); i > 0; j, i = j+1, i-1 {
		// Convert the rune at string index into an integer for mathing
		val, err := strconv.Atoi(string(numString[i-1]))
		if err != nil {
			return false
		}

		if j % 2 == 1 {
			// Double every other number, substract 9 if product is >= 9, add to sum
			product := val * 2
			if product >= 9 {
				product -= 9
			}

			sum += product
		} else {
			// Otherwise, just add number to sum
			sum += val
		}
	}

	return sum % 10 == 0
}
