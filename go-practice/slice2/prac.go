package main

import "fmt"

func main() {
	var s []int
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = a[2:4]

	a[2] = 33
	a[3] = 44
	fmt.Println("lengh of slice =", len(s))

	fmt.Println("Capacity of slice =", cap(s))
}
