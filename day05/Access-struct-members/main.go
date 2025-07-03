package main

import "fmt"

type student struct {
	name   string
	rollno int
	marks  int
}

func main() {

	var stud1 student
	var stud2 student

	//student 1
	stud1.name = "Vansh"
	stud1.rollno = 30
	stud1.marks = 95

	//student 2
	stud2.name = "aditya"
	stud2.rollno = 1
	stud2.marks = 90

	fmt.Println("Name:", stud1.name)
	fmt.Println("rollno:", stud1.rollno)
	fmt.Println("marks:", stud1.marks)

	fmt.Println("Name:", stud2.name)
	fmt.Println("rollno:", stud2.rollno)
	fmt.Println("marks:", stud2.marks)
}
