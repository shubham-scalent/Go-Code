package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]

	fmt.Printf("before s=%v\n", s)
	fmt.Printf("before a=%v\n", a)
	fmt.Printf("before len=%d,cap=%d\n", len(s), cap(s))
	fmt.Println("&a[2]==&s[0]", &a[2] == &s[0])

	s = append(s, 50, 60, 70, 80, 90, 100, 110)

	fmt.Printf("after news=%v\n", s)
	fmt.Printf("after a=%v\n", a)
	fmt.Printf("after len=%d, cap=%d\n", len(s), cap(s))
	fmt.Println("&a[2]==&s[0]", &a[2] == &s[0])
}
