package main

import "fmt"

func main() {
	myslice := []int{10, 20, 30}
	myslice = append(myslice, 40, 50)

	//append one slice to another slice
	myslice1 := append(myslice, 300)

	fmt.Println(myslice1)

	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)
}
