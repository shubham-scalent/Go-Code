package main

import "fmt"

func main() {

	a := 0x00
	b := 0x0A
	c := 0xff

	fmt.Println("&a =", &a)

	fmt.Println("&b =", &b)

	fmt.Println("&c =", &c)
}
