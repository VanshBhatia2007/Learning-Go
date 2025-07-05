package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("sample error")
	if err != nil {
		fmt.Println(err)
	}

}
