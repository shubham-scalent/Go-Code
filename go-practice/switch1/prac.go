package main

import "fmt"

func main() {

	switch letter := "m"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("letter is a vovel")
	default:
		fmt.Println("letter is not vovel")
	}

}
