package main

import "fmt"

func main() {
	today := 20
	yest := 30

	today = increment(today, yest)
	fmt.Println("total is", today)

}

func increment(today, yest int) int {
	today = today + yest
	return today
}
