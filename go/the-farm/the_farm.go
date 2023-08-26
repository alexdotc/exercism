package thefarm

import "errors"
import "fmt"

func DivideFood(calc FodderCalculator, numCows int) (float64, error) {
	fodder, err := calc.FodderAmount(numCows)
	if err != nil { return 0, err }
	fat, err := calc.FatteningFactor()
	if err != nil { return 0, err }
	return (fat * fodder) / float64(numCows), nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, numCows)
}

type InvalidCowsError struct{
	num int
	msg string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.num, err.msg)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{ numCows, "there are no negative cows" }
	}
	if numCows == 0 {
		return &InvalidCowsError{ numCows, "no cows don't need food" }
	}
	return nil
}
