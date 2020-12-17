package main

import "fmt"

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func main() {
	var n int

	for {
		num, err := fmt.Scan(&n)
		if err != nil || num == 0 {
			break
		}

		m := make([]int, n)
		dp1 := make([]int, n)
		dpr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&m[i])

			dp1[i] = 1
			for j := 0; j < i; j++ {
				if m[i] > m[j] {
					dp1[i] = max(dp1[i], dp1[j]+1)
				}
			}
		}

		for i := n - 1; i >= 0; i-- {
			dpr[i] = 1
			for j := n - 1; j >= i; j-- {
				if m[i] > m[j] {
					dpr[i] = max(dpr[i], dpr[j]+1)
				}
			}
		}

		k := 0
		for i := 0; i < n; i++ {
			if dp1[i]+dpr[i] > k+1 {
				k = dp1[i] + dpr[i] - 1
			}
		}
		fmt.Println(n - k)
	}
}
