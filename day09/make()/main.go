package main

import "fmt"

func main() {
	myslice := make([]int, 5, 10)
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	//If the capacity parameter is not defined, it will be equal to length.
	myslice2 := make([]int, 5)
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
}
