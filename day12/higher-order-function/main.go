package main

import (
	"fmt"
	"math"
)

func sphere(vol func(radius float64) float64) {
	fmt.Println("volume of sphere:", vol(5))

}

func main() {
	vol_of_sphere := func(radius float64) float64 {
		result := 4 / 3 * math.Pi * radius * radius * radius
		return result
	}
	sphere(vol_of_sphere)
}
