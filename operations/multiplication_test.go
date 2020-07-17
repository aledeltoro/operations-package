package operations

import (
	"fmt"
	"testing"
)

func TestValidateNumberSlice(t *testing.T) {

	assertNoErr := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("got an error, didn't want one")
		}
	}

	assertErr := func(t *testing.T, err error) {
		t.Helper()

		if err == nil {
			t.Errorf("didn't get an error, but wanted one")
		}
	}

	var positiveTests = []struct {
		numbers []string
	}{
		{[]string{"1", "2", "3", "4"}},
		{[]string{"2", "534", "9540", "8"}},
	}

	for _, test := range positiveTests {
		message := fmt.Sprintf("Valid number slice: %v", test.numbers)
		t.Run(message, func(t *testing.T) {
			_, err := ValidateNumberSlice(test.numbers)
			assertNoErr(t, err)
		})
	}

	var negativeTests = []struct {
		numbers []string
	}{
		{[]string{"1", "world"}},
		{[]string{"hello", "2"}},
	}

	for _, test := range negativeTests {
		message := fmt.Sprintf("Valid number slice: %v", test.numbers)
		t.Run(message, func(t *testing.T) {
			_, err := ValidateNumberSlice(test.numbers)
			assertErr(t, err)
		})
	}

}

func TestMultiplyNumbers(t *testing.T) {

	assertNoErr := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("got an error, didn't want one")
		}
	}

	assertErr := func(t *testing.T, err error) {
		t.Helper()

		if err == nil {
			t.Errorf("didn't get an error, but wanted one")
		}
	}

	var positiveTests = []struct {
		numbers []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{2, 3, 4, 5, 6}},
	}

	for _, test := range positiveTests {
		message := fmt.Sprintf("valid operation with slice: %v", test.numbers)
		t.Run(message, func(t *testing.T) {
			_, err := MultiplyNumbers(test.numbers)
			assertNoErr(t, err)
		})
	}

	var negativeTests = []struct {
		numbers []int
	}{
		{[]int{1, 2, 1<<63 - 1, 4, 5}},
		{[]int{-1 << 63, 3, 4, 5, -1 << 63}},
	}

	for _, test := range negativeTests {
		message := fmt.Sprintf("operation overflow with slice: %v", test.numbers)
		t.Run(message, func(t *testing.T) {
			_, err := MultiplyNumbers(test.numbers)
			assertErr(t, err)
		})
	}

}
