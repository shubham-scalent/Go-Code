package main

import "fmt"

func main() {
	monica := struct {
		firstname, lastname string
		salary              int
		fulltime            bool
	}{
		firstname: "Monica",
		lastname:  "Gellar",
		salary:    1200,
	}
	fmt.Println(monica)
}
