// Bob is a lackadaisical teenager
package bob

import (
	"regexp"
	"strings"
)

const (
	questionAnswer     = "Sure."
	yellAnswer         = "Whoa, chill out!"
	yellQuestionAnswer = "Calm down, I know what I'm doing!"
	emptyAnswer        = "Fine, be that way!"
	defaultAnswer      = "Whatever."
)

// Returns bob's response
func Hey(remark string) string {
	// Test for empty remark
	if len(remark) <= 0 {
		return emptyAnswer
	}

	// Strip out all non-letters
	nonletters := regexp.MustCompile(`\d`)
	remark = nonletters.ReplaceAllString(remark, "")

	// Test for question mark
	if strings.Contains(remark, "?") {
		if remark == strings.ToUpper(remark) {
			return yellQuestionAnswer
		} else {
			return questionAnswer
		}
	}

	// Test if string if remark is all uppercase
	if remark == strings.ToUpper(remark) {
		return yellAnswer
	}

	return defaultAnswer
}
