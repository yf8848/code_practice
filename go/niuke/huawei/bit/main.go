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

		fmt.Println(num2bytes(n))
	}
}

func num2bytes(n int) int {
	var max, tmp int
	last := false
	for q := n; q > 0; q = q / 2 {
		c := q % 2
		if c == 1 {
			last = true
			tmp++
		} else {
			if last && max < tmp {
				max = tmp
			}
			tmp = 0
			last = false
		}
	}
	if last && max < tmp {
		max = tmp
	}

	return max
}
