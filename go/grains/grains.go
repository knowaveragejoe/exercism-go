package grains

import (
	"errors"
	"fmt"
	"math"
)

func Square(i int) (uint64, error) {
	// Make sure we have a valid square.
	if i <= 0 || i > 64 {
		return 1, errors.New("given square does not exist on the board")
	}

	if i == 1 {
		return uint64(1), nil
	}

	return uint64(math.Pow(2, float64(i-1))), nil
}

func Total() uint64 {
	var total uint64

	for i := 0; i <= 64; i++ {
		if amount, err := Square(i); err != nil {
			fmt.Println(err)
		} else {
			total += amount
		}
	}

	return total
}
