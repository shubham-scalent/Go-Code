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
	var s Shape = Cube{4}
	c := s.(Cube)

	fmt.Println("area of c of type cube is ", s.Area())
	fmt.Println("Volume of c of type cube is ", c.Volume())
}
