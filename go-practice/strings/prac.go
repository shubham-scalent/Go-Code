package main

import "fmt"

func main() {
	var s string
	s = "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
}
