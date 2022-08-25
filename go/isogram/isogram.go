package isogram

import "strings"

// IsIsogram returns true if the given word is an isogram.
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, char := range word {

		if string(char) == "-" || string(char) == " " {
			continue
		}

		if strings.Count(word, string(char)) > 1 {
			return false
		}
	}
	return true
}
