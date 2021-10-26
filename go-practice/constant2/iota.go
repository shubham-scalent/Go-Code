package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
	)
	const a2 = iota
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)

}
