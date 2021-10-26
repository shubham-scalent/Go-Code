package main

import "fmt"

type Employee struct {
	firstname, lastname string
	salary              int
	fulltime            bool
}

func main() {
	var ross Employee

	ross.firstname = "Ross"
	ross.lastname = "Bing"
	ross.salary = 1200
	ross.fulltime = true

	fmt.Println("ross.firstname", ross.firstname)
	fmt.Println("ross.lastname", ross.lastname)
	fmt.Println("ross.salary", ross.salary)
	fmt.Println("ross.fulltime", ross.fulltime)
}
