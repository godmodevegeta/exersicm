package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, no_of_cows int) (float64, error) {
	total_fodder, err := fodderCalculator.FodderAmount(no_of_cows)
	if err != nil {
		return 0.0, err
	}
	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return total_fodder * factor / float64(no_of_cows), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, no_of_cows int) (float64, error){	
	if no_of_cows <=0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalculator, no_of_cows)

}

type InvalidCowsError struct {
	no_of_cows int
	message string
}

func (invalidCowsError *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", invalidCowsError.no_of_cows, invalidCowsError.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(no_of_cows int) error{
	if no_of_cows < 0 {
		return &InvalidCowsError{
			no_of_cows: no_of_cows,
			message: "there are no negative cows",
		}
	}	
	if no_of_cows == 0 {
		return &InvalidCowsError{
			no_of_cows: no_of_cows,
			message: "no cows don't need food",
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
