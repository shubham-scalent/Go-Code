package main

import "fmt"

func main() {
	const (
		_ = iota + 5
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)
	fmt.Printf("%v\n", snakeSpecialist)

}
