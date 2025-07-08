package main

import "fmt"

func main() {
	b := make(map[string]int)
	b["vansh"] = 10
	b["aditya"] = 20

	delete(b, "vansh")
	fmt.Println(b)
}
