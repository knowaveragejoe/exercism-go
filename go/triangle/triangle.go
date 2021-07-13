package triangle

import "math"

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Any side with <= zero length
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// Any side with NaN value
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	// Any side with Inf value
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	// Test triangle inequality theorem
	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	// All sides the same
	if a == b && b == c {
		return Equ
	}

	// Scalene - all sides are different
	if a != b && b != c && a != c {
		return Sca
	}

	// If we got this far, we know it's Isoceles
	return Iso
}
