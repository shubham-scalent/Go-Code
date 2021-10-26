package main

import "fmt"

func main() {
	// var s string = "sam"
	s := false

	fmt.Printf("%T\n", s)
	// age := 25
	var age int = 25
	fmt.Printf("%T\n", age)

	fmt.Println("His name is", s)

	fmt.Println("His age is", age, "Years")
}
