package main

import "fmt"

type Salary struct {
	basic     int
	insurance int
	allowance int
}
type Employee struct {
	firstname, lastname string
	Salary
}

func main() {
	ross := Employee{

		firstname: "Ross",
		lastname:  "Bing",
		Salary:    Salary{1100, 50, 50},
	}
	ross.basic = 1200
	ross.insurance = 100
	ross.allowance = 150

	fmt.Println("Ross's Basic Salary", ross.basic)
	fmt.Println(ross)
}
