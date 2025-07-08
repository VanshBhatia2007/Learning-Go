package main

import "fmt"

func main() {
	b := make(map[string]int)
	b["vansh"] = 10
	b["aditya"] = 20

	b["vansh"] = 50
	fmt.Println(b)
}
