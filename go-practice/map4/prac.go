package main

import "fmt"

func main() {

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	delete(age, "john")

	fmt.Println(age)

}
