package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func substract(a float32, b float32) float32 {
	return a - b
}

func main() {
	fmt.Printf("Type of function add is           %T\n", add)

	fmt.Printf("Type of function substract is     %T\n", substract)
}
