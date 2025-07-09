package main

import "fmt"

func main() {
	b := make(map[string]int)
	b["vansh"] = 10
	b["aditya"] = 20

	var a []string
	a = append(a, "vansh", "aditya")

	//no order
	for k, v := range b {
		fmt.Printf(" %v : %v ", k, v)
	}

	fmt.Println()

	for _, element := range a {
		fmt.Printf(" %v : %v ", element, b[element])
	}

}
