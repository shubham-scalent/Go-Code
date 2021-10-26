package main

import "fmt"

func endTime(timestamp string) {
	fmt.Println("Program ended at", timestamp)
}
func main() {
	time := "4 PM"

	defer endTime(time)

	time = "1 Pm"

	fmt.Println("Doing Something")
	fmt.Println("Main Finished")
	fmt.Println("time is", time)
}
