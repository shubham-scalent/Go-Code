package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[3:4]

	news := append(s, 55, 66)

	fmt.Printf("s==%v, news==%v\n", s, news)

	fmt.Printf("len==%d, cap==%d\n", len(news), cap(news))

	fmt.Printf("a==%v\n", a)

}
