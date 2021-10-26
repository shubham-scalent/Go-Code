package main

import (
	"fmt"
	"time"
)

func sqrWorker(tasks <-chan int, results chan<- int, instance int) {

	for num := range tasks {
		time.Sleep(time.Millisecond)

		fmt.Printf("[worker %v] sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}
func main() {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}
	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}
	fmt.Println("[main] wrote 5 tasks")
	close(tasks)

	for i := 0; i < 5; i++ {
		results := <-results

		fmt.Println("[main] Result ", i, ":", results)
	}
	fmt.Println("[main] main() stopped")
}
