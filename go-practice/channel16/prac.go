package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}
func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello From Service1"
}
func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello From Service2"
}
func main() {
	fmt.Println("main() strated", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response From Service1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response From Service2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("NO Response received", time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}
