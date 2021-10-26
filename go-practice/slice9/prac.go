package main

import "fmt"

func main() {

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("s[:]", s[:])

	fmt.Println("s[2:]", s[2:])

	fmt.Println("s[:4]", s[:4])

	fmt.Println("s[2:4]", s[2:4])
}
