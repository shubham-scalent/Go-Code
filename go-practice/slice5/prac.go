package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("s=", s)
	fmt.Printf("len=%d , cap=%d\n", len(s), cap(s))

	news := append(s, 7, 8)

	fmt.Println("news=", news)
	fmt.Printf("len=%d , cap=%d", len(news), cap(news))
}
