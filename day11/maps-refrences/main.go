package main

import "fmt"

func main() {
	b := make(map[string]int)
	b["vansh"] = 10
	b["aditya"] = 20
	a := b

	fmt.Println(b, a)

	a["aditya"] = 40
	fmt.Println(b, a)
}
