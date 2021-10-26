package main

import "fmt"

func main() {

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	for key, value := range age {
		fmt.Println(key, "=", value)
	}
}
