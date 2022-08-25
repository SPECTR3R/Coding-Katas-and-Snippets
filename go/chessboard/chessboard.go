package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	cRank := cb[rank]
	count := 0
	for _, val := range cRank {
		if val {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0

	if file < 0 || file > 9 {
		return count
	}

	for _, val := range cb {
		if val[file-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, rankV := range cb {
		for range rankV {
			count++
		}
	}
	return count
}

func CountOccupied(cb Chessboard) int {
	count := 0
	for rank := range cb {
		count += CountInRank(cb, rank)
	}
	return count
}
