package main

import "fmt"

func main() {
	if fruit := "apple"; fruit == "mango" {
		fmt.Println("This is a mango")
	} else if fruit == "orange" {
		fmt.Println("This is orange")
	} else if fruit == "banana" {
		fmt.Println("This is banana")
	} else {
		fmt.Println("i dont know which fruit Is this")
	}
}
