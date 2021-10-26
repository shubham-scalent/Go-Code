package main

import "fmt"

func main() {
	greetings := [...]string{
		"Good Morning",
		"Good Afternoon",
		"Good Evening",
		"Good Night",
	}
	fmt.Println(greetings)
	fmt.Println(len(greetings))

}
