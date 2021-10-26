package main

import "fmt"

func getnumber() int {
	return 20
}
func main() {
	number := 15
	switch number = getnumber(); {
	case number <= 5:
		fmt.Println("number is less than or equal to 5")
	case number > 5:
		fmt.Println("number is greater than 5")
	case number > 10:
		fmt.Println("number is greater than 10")
	case number > 15:
		fmt.Println("number is greater than 15")
	}
}
