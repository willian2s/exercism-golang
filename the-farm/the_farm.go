package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	totalFodder, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	amount := totalFodder * fatteningFactor / float64(cows)
	return amount, nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalculator, cows)
}

type InvalidCowsError struct {
	NumberOfCows  int
	CustomMessage string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.NumberOfCows, err.CustomMessage)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalidCowsError{
			NumberOfCows:  numberOfCows,
			CustomMessage: "there are no negative cows",
		}
	case numberOfCows == 0:
		return &InvalidCowsError{
			NumberOfCows:  numberOfCows,
			CustomMessage: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
