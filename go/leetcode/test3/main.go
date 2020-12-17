package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var n, m string
	var num int
	for {
		_, err := fmt.Scan(&m, &n, &num)
		if err != nil {
			break
		}

		arrM := strings.Split(m, ",")
		arrN := strings.Split(n, ",")

		numM, err := strinToInt(arrM)
		if err != nil {
			break
		}

		numN, err := strinToInt(arrN)
		if err != nil {
			break
		}

		fmt.Printf("%d\n", process(numM, numN, num))
	}

}

func process(m, n []int, num int) int {
	get := make([]int, len(n))
	for i := range n {
		get[i] = n[i] - m[i]
	}

	for len(get) > 0 && num > 0 {
		max, idx := getMax(get)
		if num >= m[idx] {
			num += max
			get[idx] = get[len(get)-1]
			get = get[:len(get)-1]
		}
	}
	return num
}

func strinToInt(strs []string) ([]int, error) {
	nums := make([]int, len(strs))
	for i, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}

	return nums, nil
}

func getMax(dp []int) (int, int) {
	max := 0
	idx := 0
	for i, val := range dp {
		if max < val {
			max = val
			idx = i
		}
	}

	return max, idx
}
