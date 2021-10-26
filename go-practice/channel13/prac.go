package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}
func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "value 1"
	chan1 <- "value 2"
	chan2 <- "value 1"
	chan2 <- "value 2"

	select {

	case res := <-chan1:
		fmt.Println("Respomce from cha1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Responce from cha2", res, time.Since(start))
	}
	fmt.Println("main() stopped", time.Since(start))
}
