package main

import "fmt"

func main() {
	myslice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//copy() function
	neededvalues := myslice[:len(myslice)-2]
	newslice := make([]int, len(neededvalues))
	copy(newslice, neededvalues)

	fmt.Println(myslice, len(myslice), cap(myslice))
	fmt.Println(newslice, len(newslice), cap(newslice))

}
