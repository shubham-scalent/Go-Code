package main

import "fmt"

func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	}
	return 1
}
func main() {
	f := getFactorial(4)
	fmt.Println(f)
}
