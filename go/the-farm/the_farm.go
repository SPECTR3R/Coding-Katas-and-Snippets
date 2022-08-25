package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	value int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", e.value)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		switch err.Error() {
		case "non-scale error":
			return 0, errors.New("non-scale error")
		case "sensor error":
			fodder *= 2
		}
	}

	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, &SillyNephewError{
			value: cows,
		}
	}

	return fodder / float64(cows), nil
}
