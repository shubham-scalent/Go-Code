package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	e.name = newName
}
func (e Employee) showSalary() {
	e.salary = 1500
	fmt.Println("salary of e =", e.salary)
}
func main() {
	e := Employee{
		name:   "Ross Galler",
		salary: 1200,
	}
	fmt.Println("e before change =", e)

	e.changeName("Monica Galler")

	(&e).showSalary()

	fmt.Println("e after change =", e)
}
