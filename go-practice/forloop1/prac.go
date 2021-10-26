package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Printf("Current Number is : %d\n", i)
		i++
		if i == 6 {
			break
		}
	}

}
