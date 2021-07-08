package hamming

import (
	"errors"
)

// Return the hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) == 0 && len(b) == 0 {
		return 0, nil
	}

	if len(a) != len(b) {
		return -1, errors.New("DNA strands are not the same length!")
	}

	hammingDistance := 0

	// Loop over A's characters and compare with B at the same index
	for i := range a {
		if a[i] != b[i] {
			hammingDistance += 1
		}
	}

	return hammingDistance, nil
}
