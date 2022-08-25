// Package scrabble computes the scrabble score of a word
package scrabble

import "strings"

var letterScore = make(map[rune]int)

func init() {
	classes := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}
	for letters, score := range classes {
		for _, letter := range letters {
			letterScore[letter] = score
		}
	}
}

func Score(word string) (score int) {
	for _, char := range strings.ToLower(word) {
		score += letterScore[char]
	}
	return
}
