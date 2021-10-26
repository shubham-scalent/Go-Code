package main

import "fmt"

func getMultiples(factor int, args ...int) []int {

	multiples := make([]int, len(args))

	for index, value := range args {
		multiples[index] = value * factor
	}
	return multiples
}
func main() {
	s := []int{10, 20, 30}

	mult1 := getMultiples(2, s...)

	mult2 := getMultiples(3, 1, 2, 3, 4)

	mult3 := getMultiples(4, 2, 4, 6, 8)

	fmt.Println(mult1)
	fmt.Println(mult2)
	fmt.Println(mult3)

}
