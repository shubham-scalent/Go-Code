package main

import "fmt"

func main() {

	a := 1
	p := &a
	*p = 2
	fmt.Printf("a = %v\n", a)
	fmt.Printf("data at %v is %v\n", p, *p)
}
