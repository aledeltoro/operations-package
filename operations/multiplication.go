package operations

import (
	"errors"
	"strconv"
)

var (
	errInvalidType = errors.New("error: invalid variable type")
)

// ValidateNumberSlice checks if the given arguments are integer values
func ValidateNumberSlice(numbers []string) ([]int, error) {

	var validatedNumbers []int

	for _, number := range numbers {

		validNum, err := strconv.Atoi(number)
		if err != nil {
			return nil, errInvalidType
		}

		validatedNumbers = append(validatedNumbers, validNum)

	}

	return validatedNumbers, nil

}

// MultiplyNumbers multiplies the slice of validated numbers
func MultiplyNumbers(numbers []int) (int, error) {

	var result int = 1

	for _, number := range numbers {
		result = result * number
	}

	return result, nil

}
