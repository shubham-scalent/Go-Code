package main

import "fmt"

type Contact struct {
	phone, address string
}
type Employee struct {
	name   string
	salary int
	Contact
}

func (c *Contact) changePhone(newPhone string) {
	c.phone = newPhone
}
func main() {
	e := Employee{
		name:   "Ross aller",
		salary: 1300,
		Contact: Contact{
			phone:   "9876543210",
			address: "New Delhi,India",
		},
	}
	fmt.Println("e before phone change =", e)

	e.Contact.changePhone("987654329")

	fmt.Println("e after phone change =", e)
}
