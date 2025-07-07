package main

import "fmt"

func main() {
	myslice := []int{10, 20, 30}
	myslice[1] = 40

	fmt.Println(myslice[0])
	fmt.Println(myslice[1])
	fmt.Println(myslice[2])
}
