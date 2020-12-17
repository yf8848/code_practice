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

	len1, len2 := len(str1), len(str2)
	dp := make([][]int, len1)

	max, pos := 0, 0
	for i := 0; i < len1; i++ {
		tmp := make([]int, len2)

		for j := 0; j < len2; j++ {
			if str1[i] == str2[j] {
				if i == 0 || j == 0 {
					tmp[j] = 1
				} else {
					tmp[j] = dp[i-1][j-1] + 1
				}

				if max < tmp[j] {
					max = tmp[j]
					pos = i
				}
			}
		}
		dp[i] = tmp
	}

	if max == 0 {
		return "-1"
	}
	return str1[pos-max+1 : pos+1]
}
