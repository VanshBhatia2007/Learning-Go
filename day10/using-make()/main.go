package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["vansh"] = "v"
	a["aditya"] = "a"

	b := make(map[string]int)
	b["vansh"] = 10
	b["aditya"] = 20

	fmt.Println(a)
	fmt.Println(b)
}
