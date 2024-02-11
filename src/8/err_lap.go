package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not.txt")
	if err != nil {
		fmt.Println(err)
		if wrapperErr := errors.Unwrap(err); wrapperErr != nil {
			fmt.Println("in main, weapperErr: ", wrapperErr)
		}
	}
}
