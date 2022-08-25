package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming distance between two strings.
// The Hamming distance is the number of positions at which the
// corresponding symbols are different.
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("left and right strands must be of equal length")
	}
	count := 0
	for pos, char := range a {
		if char != rune(b[pos]) {
			count++
		}
	}
	return count, nil
}
