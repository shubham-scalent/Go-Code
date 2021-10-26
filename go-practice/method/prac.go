package main

import "fmt"

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
func main() {
	value1 := data(23)
	value2 := data(20)

	result := value1.multiply(value2)
	fmt.Println("final result", result)
}
