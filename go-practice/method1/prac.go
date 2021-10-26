package main

import (
	"fmt"
	"math"
)

type rect struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func main() {

	rect := rect{5.0, 4.0}
	circ := circle{5.0}

	fmt.Printf("Area of Rectangle = %0.2f\n", rect.Area())
	fmt.Printf("Area of Circle = %0.2f\n", circ.Area())
}
