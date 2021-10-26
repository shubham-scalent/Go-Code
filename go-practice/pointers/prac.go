package main

import "fmt"

func main() {

	a := 0x00
	b := 0x0A
	c := 0xff

	fmt.Printf("variable a of type %T with value %v in hexa is %X\n", a, a, a)

	fmt.Printf("variable a of type %T with value %v in hexa is %X\n", b, b, b)

	fmt.Printf("variable a of type %T with value %v in hexa is %X\n", c, c, c)
}
