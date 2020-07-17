package main

import (
	"fmt"
	"log"
	"os"

	multiply "exam/operations"
)

func main() {

	arguments := os.Args[1:]

	validatedNumbers, err := multiply.ValidateNumberSlice(arguments)
	if err != nil {
		log.Fatal(err)
	}

	result, err := multiply.MultiplyNumbers(validatedNumbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
