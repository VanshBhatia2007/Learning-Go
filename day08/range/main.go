package main

import "fmt"

func main() {
	rvariable := []string{"h", "hel", "hello"}

	for i, j := range rvariable {
		fmt.Println(i, j)
	}
}
