package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		arr := make([]int, n)
		var num, sum5, sum3 int

		idx := 0
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			if num%5 == 0 {
				sum5 += num
			} else if num%3 == 0 {
				sum3 += num
			} else {
				arr[idx] = num
				idx++
			}
		}

		fmt.Println(equal(sum3, sum5, arr))
	}
}

func equal(sum3, sum5 int, arr []int) bool {
	if len(arr) == 0 {
		return sum3 == sum5
	} else {
		return equal(sum3+arr[0], sum5, arr[1:]) || equal(sum3, sum5+arr[0], arr[1:])
	}
}
