package main

import "fmt"

func main() {

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}
	fmt.Println(age["mina"])
	fmt.Println(age["john"])
	fmt.Println(age["mike"])
	fmt.Println(age["jordan"])

	fmt.Println("len(age)=", len(age))
}
