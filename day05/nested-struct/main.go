package main

import "fmt"

type student struct {
	name  string
	marks int
}

type result struct {
	english student
	maths   student
}

func main() {
	final := result{
		english: student{name: "vansh", marks: 90},
		maths:   student{name: "vansh", marks: 99},
	}
	fmt.Println(final.english.marks)
}
