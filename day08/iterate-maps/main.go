package main

import "fmt"

func main() {
	mmap := map[int]string{
		01: "you",
		02: "are",
		03: "noob",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}
}
