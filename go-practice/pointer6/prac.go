package main

import "fmt"

func changeValue(p *int) {

	*p = 2
}

func main() {

	a := 1

	changeValue(&a)

	fmt.Printf("a = %v\n", a)
}
