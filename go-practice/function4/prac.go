package main

import "fmt"

func add(a, b int) int32 {
	return int32(a + b)
}
func main() {
	result := add(5, 9)
	fmt.Println(result)
}
