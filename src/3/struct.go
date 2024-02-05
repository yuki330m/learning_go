package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	bob := person{}
	julia := person{
		name: "ジュリア",
		age:  40,
	}
	fmt.Println(fred.name)
	fmt.Println(bob)
	fmt.Println(julia)

	var person1 struct {
		name string
		age  int
		pet  string
	}

	person1.name = "bob"
	person1.age = 40
	person1.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "poti",
		kind: "dog",
	}

	fmt.Println(person1)
	fmt.Println(pet)
}
