package main

import "fmt"

func split(sum int) (x int, y int) {
	x = sum - 10 + 20
	y = x + 50
	return

}

func main() {
	fmt.Println(split(10))
}
