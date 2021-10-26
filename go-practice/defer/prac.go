package main

import "fmt"

func sayDone() {
	fmt.Println("I am Done")
}
func main() {
	fmt.Println("Main Started")

	defer sayDone()

	fmt.Println("Main Finished")
}
