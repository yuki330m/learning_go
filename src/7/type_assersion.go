package main

import "fmt"

type Myint int

func main() {
	var i any
	var mine Myint = 20
	i = mine
	i2, ok := i.(int)
	fmt.Println(i2)
	fmt.Println(ok)
}

func doTypeSwitch(i any) {
	switch j := i.(type) {
	case nil:
	case int:
	case Myint:
	}
}
