package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	e.name = newName
}
func main() {
	e := Employee{
		name:   "Ross Galler",
		salary: 1200,
	}
	fmt.Println("e befor name change =", e)

	e.changeName("Monica Galler")

	fmt.Println("e after name change =", e)
}
