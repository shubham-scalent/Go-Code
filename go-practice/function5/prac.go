package main

import "fmt"

func addMult(a, b int) (add, mult int) {
	add = a + b
	mult = a * b

	return
}

func main() {
	add, mult := addMult(2, 5)
	fmt.Println(add, mult)
}
