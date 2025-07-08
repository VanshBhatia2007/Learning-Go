package main

import "fmt"

func main() {
	b := make(map[string]int)
	b["vansh"] = 10
	b["aditya"] = 20

	val1, ok1 := b["vansh"]
	_, ok2 := b["aditya"]
	val3, ok3 := b["hello"]

	fmt.Println(val1, ok1)
	fmt.Println(ok2)
	fmt.Println(val3, ok3)
}
