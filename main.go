package main

import (
	"fmt"
	"os"
	"log"

	multiply "exam/operations"
	
)

func main() {

	numbers := os.Args[1:]

	validatedNumbers, err := multiply.ValidateNumberSlice(numbers)
	if err != nil {
		log.Fatal(err)
	}

	result := multiply.MultiplyNumbers(validatedNumbers)

	fmt.Println(result)

}


