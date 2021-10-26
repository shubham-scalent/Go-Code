package main

import "fmt"

func main() {

	var s []int
	fmt.Println("s == nil", s == nil)

	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = a[2:4]
	fmt.Println("s == nil", s == nil, "and s =", s)
}
