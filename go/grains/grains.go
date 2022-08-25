package grains

import "errors"

// Square returns the number of grains on the square. Error if input is not a valid square.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid square")
	}
	result := 1 << uint64(number-1)
	return uint64(result), nil
}

func Total() uint64 {
	total := uint64(0)
	for i := 1; i <= 64; i++ {
		val, _ := Square(i)
		total += val
	}
	return uint64(total)
}
