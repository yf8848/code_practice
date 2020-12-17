package main

import (
	"fmt"
)

/**
 * https://www.nowcoder.com/practice/f9c6f980eeec43ef85be20755ddbeaf4?tpId=37&&tqId=21239&rp=1&ru=/ta/huawei&qru=/ta/huawei/question-ranking
 */

const SZ = 61

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	N := 0
	m := 0

	_, err := fmt.Scan(&N, &m)
	if err != nil {
		panic(err)
	}

	//价格是10的倍数，降低维度
	N /= 10
	price := make([][]int, SZ)
	for i := 0; i < SZ; i++ {
		price[i] = make([]int, 3)
	}

	wight := make([][]int, SZ)
	for i := 0; i < SZ; i++ {
		wight[i] = make([]int, 3)
	}

	for i := 1; i <= m; i++ {
		a, b, c := 0, 0, 0
		fmt.Scan(&a, &b, &c)
		a /= 10
		b *= a
		if c == 0 {
			price[i][0] = a
			wight[i][0] = b
		} else {
			if price[c][1] == 0 {
				price[c][1] = a
				wight[c][1] = b
			} else {
				price[c][2] = a
				wight[c][2] = b
			}
		}
	}

	// 分组背包
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= N; j++ {
			a := price[i][0]
			x := wight[i][0]
			b := price[i][1]
			y := wight[i][1]
			c := price[i][2]
			z := wight[i][2]

			if j >= a {
				dp[i][j] = max(dp[i-1][j-a]+x, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}

			if j >= a+b {
				dp[i][j] = max(dp[i-1][j-a-b]+x+y, dp[i][j])
			}
			if j >= a+c {
				dp[i][j] = max(dp[i-1][j-a-c]+x+z, dp[i][j])
			}
			if j >= a+b+c {
				dp[i][j] = max(dp[i-1][j-a-b-c]+x+y+z, dp[i][j])
			}
		}
	}
	fmt.Println(dp[m][N] * 10)
}
