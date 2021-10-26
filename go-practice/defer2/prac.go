package main

import "fmt"

func greet(message string) {
	fmt.Println("greeting :", message)
}

func main() {

	fmt.Println("call one")
	defer greet("greet one")

	fmt.Println("call two")
	defer greet("greet two")

	fmt.Println("call three")
	defer greet("greet three")
}
