package main

import (
	"fmt"
)

func main() {
	var n int
    for{
		//num,err:=fmt.Scanf("%d", &n)
		//!!!Scanf无法 AC？？？
		num,err:=fmt.Scan(&n)
        if num==0||err!=nil{
            return
        }
        nums := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Scanf("%d", &nums[i])
        }
        fmt.Printf("%d\n", dlc(nums))
    }
}

func dlc(nums []int) int {
	max := 0
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = 1
	}
	for i, n := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < n && res[i] < res[j]+1 {
				res[i] = res[j] + 1
			}
		}
		if max < res[i] {
			max = res[i]
		}
	}
	return max
}
