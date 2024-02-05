package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while文
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	evenvals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenvals {
		fmt.Println(i, v)
	}

outer:
	for {
		for {
			if rand.Intn(4) == 1 {
				break outer
			}
			fmt.Println("inloop")
			break
		}
		fmt.Println("outloop")
	}
	fmt.Println("finish")
}
