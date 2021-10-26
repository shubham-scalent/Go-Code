package main

import "fmt"

func main() {

	age := map[bool]string{
		true:  "Yes",
		false: "No",
	}
	for key, value := range age {
		fmt.Println(key, "=", value)
	}
}
