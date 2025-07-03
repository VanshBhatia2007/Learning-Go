package main

import "fmt"

type Student struct {
	string
	int
}

func main() {
	student := Student{"vansh", 30}
	fmt.Println(student.string)
	fmt.Println(student.int)
}
