package main

import "fmt"

func main() {

	s1 := make([]int, 2, 4)
	s2 := []int{1, 2, 3}

	fmt.Printf("before s1=%v,s2=%v\n", s1, s2)

	n1 := copy(s1, s2)

	fmt.Printf("after n1=%v s1=%v,s2=%v\n", n1, s1, s2)

	fmt.Println("after s1==nil", s1 == nil)
}
