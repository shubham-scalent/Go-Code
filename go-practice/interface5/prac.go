package main

import "fmt"

type Shape interface {
	Area() float64
}
type Object interface {
	Volume() float64
}
type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c

	fmt.Println("volume of s of interface type shape is ", s.Area())
	fmt.Println("area of o of interface type object is", o.Volume())
}
