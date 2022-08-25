package luhn

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 2 {
		return false
	}

	parity := len(input) % 2
	sum := 0
	for i, digit := range input {
		val, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		if i%2 == parity {
			val = val * 2
			if val > 9 {
				val = val - 9
			}
		}

		sum += val
	}

	return (sum % 10) == 0
}
