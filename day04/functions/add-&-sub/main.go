package main

import "fmt"

//add
func add(x int, y int) int {
	return x + y

}

///sub
func sub(x int, y int) int {
	return x - y

}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(sub(20, 10))
}
