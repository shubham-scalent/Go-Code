package main

import "fmt"

func service(c chan string) {
	c <- "response"
}
func main() {
	fmt.Println("main() started")

	var chan1 chan string

	go service(chan1)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	default:
		fmt.Println("No Response")
	}
	fmt.Println("main() stopped")
}
