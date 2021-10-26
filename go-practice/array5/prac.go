package main

import "fmt"

func main() {
	a := [...][2]int{{1, 2}, {2, 3}, {7, 2}}

	fmt.Printf("Array is %v and type of array elements is %T\n", a, a[0])

}
