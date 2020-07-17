package operations

import (
	"errors"
	"math"
	"strconv"
)

var (
	errInvalidType = errors.New("error: invalid variable type")
	errIntOverflow = errors.New("error: data type overflow")
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

		overflow := checkOverflow(result, number)
		if overflow {
			return 0, errIntOverflow
		}

		result = result * number
	}

	return result, nil

}

func checkOverflow(num1, num2 int) bool {

	if num1 > math.MaxInt32-num2 {
		return true
	}

	if num1 < math.MinInt32-num2 {
		return true
	}

	return false

}
