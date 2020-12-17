package main

import (
	"fmt"
)

func main() {
	var str1, str2 string

	for {
		n, err := fmt.Scan(&str1, &str2)
		if err != nil || n == 0 {
			break
		}

		fmt.Println(LCS(str1, str2))
	}
}

func LCS(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return "-1"
	}

	len1 := len(str1)
	len2 := len(str2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	res := ""
	for i, j := len1, len2; i > 0 && j > 0; {
		if str1[i-1] == str2[j-1] {
			res = string(str1[i-1]) + res
			i--
			j--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}
	if len(res) == 0 {
		return "-1"
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
