package main

import "fmt"

type Shape interface {
	Area() float64
}
type Object interface {
	Volume() float64
}
type Material interface {
	Shape
	Object
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
	var m Material = c
	var s Shape = c
	var o Object = c

	fmt.Printf("Dynamic type and value of interface m of state type Material is %T and %v\n", m, m)

	fmt.Printf("Dynamic type and value of interface m of state type Shape is %T and %v\n", s, s)

	fmt.Printf("Dynamic type and value of interface m of state type Object is %T and %v\n", o, o)
}
