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

func calculatearea(shape interface{}) {
	switch s := shape.(type) {
	case circle:
		fmt.Println("area of circle", s.area())
	case rectangle:
		fmt.Println("area of rectangle", s.area())
	default:
		fmt.Println("unknown shape")
	}
}
func main() {
	c := circle{radius: 5}
	r := rectangle{length: 4, width: 3}

	calculatearea(c)
	calculatearea(r)
}
