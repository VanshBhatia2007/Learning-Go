package main

import "fmt"

func main() {
	arr1 := [7]int{10, 20, 30, 40, 50, 60}
	myslice := arr1[2:5]
	fmt.Println(myslice)
	fmt.Println(cap(myslice))
	fmt.Println(len(myslice))
}
