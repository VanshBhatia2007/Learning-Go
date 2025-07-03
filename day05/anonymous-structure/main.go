package main

import "fmt"

type Student struct {
	details struct {
		name   string
		rollno int
	}
	cgpa float64
}

func main() {
	student := Student{
		details: struct {
			name   string
			rollno int
		}{
			name:   "vansh",
			rollno: 30,
		},
		cgpa: 9.9,
	}
	fmt.Println(student.cgpa)
}
