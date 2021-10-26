package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.width
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}
func main() {
	r := Rect{5.0, 4.0}
	var s Shape = &r
	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Println("Area of Rectangke is =", area)
	fmt.Println("Perimeter of Rectangle is =", perimeter)
}
