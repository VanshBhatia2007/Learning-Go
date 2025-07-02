package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x

}

func main() {

	a, b := swap("hello", "brother")
	fmt.Println(a, b)
	fmt.Println(swap("help", "me"))

}
