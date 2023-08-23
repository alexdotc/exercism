package chessboard

type File [8]bool // array, not slice
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f := cb[file]
	count := 0
	for _, occ := range f {
		if occ == true { count++ }
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 { return 0 }
	count := 0
	for _, f := range cb {
		if f[rank-1] == true { count++ }
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, f := range cb {
		count += len(f)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, f := range cb {
		for _, v := range f {
			if v == true { count++ }
		}
	}
	return count
}
