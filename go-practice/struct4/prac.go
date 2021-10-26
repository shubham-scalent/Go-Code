package main

import "fmt"

type Employee struct {
	firstname, lastname string
	salary              int
	fulltime            bool
}

func main() {
	ross := &Employee{

		firstname: "Ross",
		lastname:  "Bing",
		salary:    1200,
		fulltime:  true,
	}
	fmt.Println("firstname", ross.firstname)
}
