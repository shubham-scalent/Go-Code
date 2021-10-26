package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]

	fmt.Println("address of a[2]", &a[2])

	fmt.Println("address of s[0]", &s[0])

	fmt.Println("&a[2] == &s[0]", &a[2] == &s[0])
}
