package main

import (
	"fmt"
	"os"
	"log"

	multiply "exam/operations"
	
)

func main() {

	arguments := os.Args[1:]

	validatedNumbers, err := multiply.ValidateNumberSlice(arguments)
	if err != nil {
		log.Fatal(err)
	}

	result := multiply.MultiplyNumbers(validatedNumbers)

	fmt.Println(result)

}


