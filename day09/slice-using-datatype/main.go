package main

import "fmt"

func main() {
	//empty slice
	myslice := []int{}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	myslice2 := []string{"hello", "how", "are", "you"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

}
