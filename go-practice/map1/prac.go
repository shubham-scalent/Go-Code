package main

import "fmt"

func main() {

	age := make(map[string]int)

	age["mina"] = 28

	age["john"] = 32

	age["mike"] = 55

	fmt.Println("age of john is =", age["john"])
}
