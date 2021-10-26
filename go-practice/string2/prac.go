package main

import "fmt"

func main() {
	s := "Hello"
	for index, char := range s {
		fmt.Printf("charcter at index %d is %c\n", index, char)
	}
}
