package main

import "fmt"

func getMultiples(factor int, args ...int) []int {
	for index, value := range args {
		args[index] = value * factor
	}
	return args
}
func main() {
	s := []int{10, 20, 30}

	mult1 := getMultiples(2, s...)

	fmt.Println(s, mult1)
}
