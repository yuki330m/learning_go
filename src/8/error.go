package main

import (
	// "errors"
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		// return 0, errors.New("処理対象は偶数のみです")
		return 0, fmt.Errorf("%dは偶数ではありません", i)

	}

	return i * 2, nil
}

func main() {
	i := 20
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%dの2倍: %d\n", i, double)
}
