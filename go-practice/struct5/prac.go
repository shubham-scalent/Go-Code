package main

import "fmt"

type Data struct {
	string
	int
	bool
}

func main() {

	sample1 := Data{"Monday", 1200, true}
	sample1.bool = false

	fmt.Println(sample1.string, sample1.int, sample1.bool)
}
