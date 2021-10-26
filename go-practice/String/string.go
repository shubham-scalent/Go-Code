package main

import "fmt"

func main() {
	s := "This Is a String"
	b := []byte(s)
	fmt.Printf("%v ,%T\n", b, b)
}
