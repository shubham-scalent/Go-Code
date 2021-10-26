package main

import "fmt"

type Contact struct {
	phone, adress string
}
type Employee struct {
	name    string
	salary  int
	contact Contact
}

func (e *Employee) changePhone(newPhone string) {
	e.contact.phone = newPhone
}
func main() {

	e := Employee{
		name:    "Ross Galler",
		salary:  1200,
		contact: Contact{"9876543210", "New Delhi"},
	}
	fmt.Println("e before phone change =", e)

	e.changePhone("987654329")

	fmt.Println("e after phone change =", e)
}
