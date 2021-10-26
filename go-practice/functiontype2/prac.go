package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

type CalcFunc func(int, int) int

func calc(a, b int, f CalcFunc) int {
	r := f(a, b)
	return r
}
func main() {
	addResult := calc(5, 3, add)
	subResult := calc(5, 3, sub)

	fmt.Println("5+3 =", addResult)
	fmt.Println("5-3 =", subResult)
}
