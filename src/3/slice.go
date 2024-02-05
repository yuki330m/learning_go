package main

import "fmt"

func main() {
	// [n],[...]は配列[]はスライス
	var x = []int{10, 20, 30}
	fmt.Println(x)
	y := append(x, 40, 50)
	fmt.Println(y)

	z := make([]int, 5)
	fmt.Println(z)
	za := append(z, 10)
	fmt.Println(za)
	zb := make([]int, 5, 10)
	fmt.Println(zb)

	a := []int{1, 2, 3, 4}
	b := a[:2]
	c := a[1:]
	d := a[1:3]
	e := a[:]
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
