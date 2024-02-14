package main

import (
	"fmt"

	"example.co.jp/package_example/formatter"
	"example.co.jp/package_example/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
