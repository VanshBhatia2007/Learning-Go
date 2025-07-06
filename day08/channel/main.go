package main

import "fmt"

func main() {
	chn1 := make(chan int)
	go func() {
		chn1 <- 100
		chn1 <- 1000
		chn1 <- 10000
		close(chn1)
	}()
	for i := range chn1 {
		fmt.Println(i)
	}
}
