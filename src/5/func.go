package main

import (
	"errors"
	"fmt"
	"os"
)

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, errors.New("0での除算はできません")
	}

	result, remainder = numerator/denominator, numerator%denominator

	return result, remainder, nil
}

func main() {
	result, remainder, err := divAndRemainder(2, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result, remainder)
}
