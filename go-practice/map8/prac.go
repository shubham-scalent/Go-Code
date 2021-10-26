package main

import "fmt"

func main() {
	ages := make(map[string]int)

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	for key, value := range age {
		ages[key] = value
	}
	delete(ages, "john")

	fmt.Println("age", age)

	fmt.Println("ages", ages)
}
