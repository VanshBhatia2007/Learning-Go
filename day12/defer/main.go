package main

import "fmt"

func mul(a, b int) int {
	result := a * b
	fmt.Println("result", result)
	return 0
}

func main() {

	mul(10, 20)

	defer mul(20, 30)
}
