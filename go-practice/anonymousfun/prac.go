package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}(3, 6)

	fmt.Println("5+3 =", sum)
}
