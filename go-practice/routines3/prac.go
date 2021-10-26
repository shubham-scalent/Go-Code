package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main execution started")
	go func() {
		fmt.Println("Hello World")
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stopped")
}
