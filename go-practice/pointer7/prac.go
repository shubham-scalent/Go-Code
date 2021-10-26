package main

import "fmt"

func changeValues(p *[4]int) {

	(*p)[0] *= 2
	(*p)[1] *= 3
	(*p)[2] *= 4
}
func main() {

	a := [4]int{1, 2, 3}

	changeValues(&a)

	fmt.Printf("a = %v\n", a)
}
