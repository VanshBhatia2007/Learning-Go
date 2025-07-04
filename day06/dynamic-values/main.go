package main

import "fmt"

type Shape interface {
	Area() float64
}

func main() {
	var s Shape
	fmt.Println("value of s:", s)
	fmt.Printf("type of s %t\n :", s)
}
