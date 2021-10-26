package main

import "fmt"

func main() {
	a := 1
	p := &a

	fmt.Printf("Pointer p of type %T with value %v\n", p, p)
	fmt.Printf("data at %v is %v\n", p, *p)
}
