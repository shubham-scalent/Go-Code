package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	for index := 0; index < len(a); index++ {

		fmt.Printf("a[%d] = %d\n", index, a[index])

	}

}
