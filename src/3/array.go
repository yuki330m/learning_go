package main

import "fmt"

func main() {
	var x [3]int
	var y = [3]int{10, 20, 30}
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
