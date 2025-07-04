package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, width float64
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	var s shape
	s = circle{radius: 10}
	fmt.Println("area of circle:", s.area())
	fmt.Println("perimeter of circle:", s.perimeter())

	s = rectangle{length: 10, width: 20}
	fmt.Println("area of rectangle:", s.area())
	fmt.Println("perimeter of rectangle:", s.perimeter())
}
