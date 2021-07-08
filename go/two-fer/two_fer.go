package twofer

import "fmt"

// Returns a ShareWith string depending on input
func ShareWith(name string) string {
	template := "One for %v, one for me."

	if len(name) <= 0 {
		name = "you"
	}

	return fmt.Sprintf(template, name)
}
