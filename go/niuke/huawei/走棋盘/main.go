package main

import "fmt"

func main() {
	var n, m int
	for {
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			return
		}

		fmt.Println(steps(0, 0, n, m))
	}
}

func steps(x, y, n, m int) int {
	if x < n && y < m {
		return steps(x+1, y, n, m) + steps(x, y+1, n, m)
	} else if x < n && y == m {
		return steps(x+1, y, n, m)
	} else if x == n && y < m {
		return steps(x, y+1, n, m)
	} else {
		return 1
	}
}
