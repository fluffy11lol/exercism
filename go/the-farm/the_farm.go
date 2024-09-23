package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, n int) (float64, error) {
	total, err := fc.FodderAmount(n)
	if err != nil {
		return 0, err
	}
	totalF, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return total * totalF / float64(n), nil

}

func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
	if n > 0 {
		return DivideFood(fc, n)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(n int) error {
	if n < 0 {
		return &InvalidCowsError{
			cows:    n,
			message: "there are no negative cows",
		}
	} else if n == 0 {
		return &InvalidCowsError{
			cows:    0,
			message: "no cows don't need food",
		}
	} else {
		return nil
	}
}
