package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Printf("%.1f\n", getCubeRoot(input))
}

func getCubeRoot(input float64) float64 {
	start := 0.0
	end := input
	mid := (start + end) / 2

	for math.Abs(mid*mid*mid-input) > 1e-5 {
		if mid*mid*mid > input {
			end = mid
		} else {
			start = mid
		}
		mid = (end + start) / 2
	}
	return mid
}
