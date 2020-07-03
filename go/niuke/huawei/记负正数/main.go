package main

import (
	"fmt"
)

func main() {
	var n, t int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		var sum, num, na int
		for i := 0; i < n; i++ {
			fmt.Scan(&t)
			if t < 0 {
				num++
			} else if t > 0 {
				sum += t
				na++
			}
		}
		fmt.Printf("%d %.1f\n", num, float64(sum)/float64(na))
	}
}
