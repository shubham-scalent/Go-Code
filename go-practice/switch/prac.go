package main

import "fmt"

func main() {
	switch i := 1 + 5; i {
	case 1, 3, 5:
		fmt.Println("One , Three and Five")
	case 2, 4, 6:
		fmt.Println("Two , four and Six")
	default:
		fmt.Println("Another Number")
	}
}
