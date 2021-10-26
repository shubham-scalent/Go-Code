package main

import "fmt"

func main() {

	s1 := make([]int, 0, 10)
	s2 := []int{1, 2, 3}

	fmt.Println("before s1 ,s2 =", s1, s2)

	s1 = append(s1, s2...)

	fmt.Println("after s1 =", s1)
}
