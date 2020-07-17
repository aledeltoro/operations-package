package operations

import (
	"fmt"
	"testing"
)

func TestValidateNumberSlice(t *testing.T) {

	assertNoErr := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("got an error, didin't want one")
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
		{[]string{"2","534","9540","8"}},
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
