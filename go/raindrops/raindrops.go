package raindrops

import "fmt"

func Convert(n int) string {
	sound := ""
	if n%3 == 0 {
		sound += "Pling"
	}
	if n%5 == 0 {
		sound += "Plang"
	}
	if n%7 == 0 {
		sound += "Plong"
	}

	// if no factors, return the number itself
	if len(sound) <= 0 {
		sound = fmt.Sprint(n)
	}

	return sound
}
