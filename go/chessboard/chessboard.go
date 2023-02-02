package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var squares int
	for _, file := range cb[file] {
		if file == true {
			squares += 1
		}
	}

	return squares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var squares int

	if rank <= 0 || rank >= 9 {
		return 0
	}

	// Loop over Files
	for _, v := range cb {
		if v[rank-1] {
			squares += 1
		}
	}

	return squares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var total int
	for i := range cb {
		total += CountInFile(cb, i)
	}

	return total
}
