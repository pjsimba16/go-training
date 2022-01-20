package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var count int
	for _, full := range cb[rank] {
		if full {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var count int
	if (file < 1) || (file > 8) {
		return 0
	} else {
		for _, rank := range cb {
			if rank[file-1] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var count int
	for _, rank := range cb {
		count += len(rank)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var count int
	for key := range cb {
		count += CountInRank(cb, key)
	}
	return count
}
