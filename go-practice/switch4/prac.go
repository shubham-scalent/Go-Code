package main

import "fmt"

func main() {
	a := 10
	b := 5
	c := 3
	switch c {
	case 1:
		c = a + b
		fmt.Printf("Addition is : %d\n", c)
		fallthrough
	case 2:
		c = a - b
		fmt.Printf("Substraction is : %d\n", c)
		fallthrough
	case 3:
		c = a * b
		fmt.Printf("Multiplicaion is : %d\n", c)
		fallthrough
	case 4:
		c = a / b
		fmt.Printf("Division is : %d\n", c)
	default:
		fmt.Println("Invalid Value")

	}
}
