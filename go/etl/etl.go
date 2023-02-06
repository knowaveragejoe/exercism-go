package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	scores := make(map[string]int)

	for i, l := range in {
		for _, c := range l {
			scores[strings.ToLower(c)] = i
		}
	}

	return scores
}
