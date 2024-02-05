package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["ライターズ"] = 1
	totalWins["ナイツ"] = 2
	fmt.Println(totalWins["ライターズ"])
	fmt.Println(totalWins["ナイツ"])
	totalWins["ミュージシャンズ"]++
	fmt.Println(totalWins["ミュージシャンズ"])
	totalWins["ナイツ"] = 3
	fmt.Println(totalWins["ナイツ"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}
