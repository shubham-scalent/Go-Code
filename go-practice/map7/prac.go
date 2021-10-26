package main

import "fmt"

func main() {
	var ages map[string]int

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}
	ages = age

	delete(ages, "john")
	fmt.Println("age", age)
	fmt.Println("ages", ages)
}
