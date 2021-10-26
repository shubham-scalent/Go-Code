package main

import "fmt"

func main() {

	c := make(chan int, 3)

	c <- 1
	c <- 2

	fmt.Printf("length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))
}
