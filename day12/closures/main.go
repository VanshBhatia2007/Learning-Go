package main

import "fmt"

func newcount() func() int {
	gfg := 0
	return func() int {
		gfg += 1
		return gfg
	}

}

func main() {
	counter := newcount()

	fmt.Println(counter())
	fmt.Println(counter())
}
