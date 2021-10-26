// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func numbers(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	fmt.Println("main() started")

	go numbers("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
