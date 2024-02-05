package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "authlopologist", "タコの足は8本でイカの足は1本"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("[%s]の文字数は%dで短い\n", word, size)
		case 5:
			fmt.Printf("[%s]の文字数は%dでちょうどいい\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("[%s]の文字数は%dでとても長い", word, size)
			if n := len(word); size < n {
				fmt.Printf("%dバイトもある\n", n)
			} else {
				fmt.Println()
			}
		}
	}
}
