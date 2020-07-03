package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		var sum int
		for i := 0; i < n; i++ {
			if isSelfNum(i) {
				sum++
			}
		}
		fmt.Printf("%d\n", sum)
	}
}

func isSelfNum(num int) bool {
	res := num * num
	resStr := strconv.Itoa(res)
	numStr := strconv.Itoa(num)
	return strings.HasSuffix(resStr, numStr)
}
