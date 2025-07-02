package main

import "fmt"

func main() {
	a, _ := myname()
	fmt.Println("my name is", a)

}

func myname() (string, string) {
	return "vansh", "bhatia"
}
