package main

import "fmt"

func main() {
	var m, n int
	for {
		num, err := fmt.Scan(&m, &n)
		if err != nil || num == 0 {
			break
		}

		if m < 0 || n < 1 {
			fmt.Println(-1)
			continue
		}
		fmt.Println(fn(m, n))
	}
}

func fn(m, n int) int {
	if n < 0 || m < 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}
	return fn(m, n-1) + fn(m-n, n)
}
