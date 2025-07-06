package main

import "fmt"

func main() {
	h := [2]string{"hello", "bye"}
	name := [2]string{"vansh", "aditya"}
	for i := 0; i < len(h); i++ {
		for j := 0; j < len(name); j++ {
			fmt.Println(h[i], name[j])
		}
	}
}
