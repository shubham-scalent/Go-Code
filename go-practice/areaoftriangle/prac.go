package main

import "fmt"

func main() {
	var radius float64
	var PI float64 = 3.14
	var area float64
	var ci float64

	fmt.Print("Enter radius of circle :")
	fmt.Scan(&radius)

	area = PI * radius * radius
	fmt.Println("area of circle is :", area)

	ci = 2 * PI * radius
	fmt.Println("Circumference of circle is :", ci)
}
