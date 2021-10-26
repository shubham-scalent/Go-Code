package main

import "fmt"

type Employee struct {
	firstname, lastname string
	salary              int
	fulltime            bool
}

func main() {

	Ross := Employee{
		firstname: "ross",
		lastname:  "Bing",
		salary:    1200,
		fulltime:  true,
	}
	fmt.Println(Ross)
}
