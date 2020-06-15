package main

import (
	"fmt"
	"io"
)

func main() {
	var num, num2, sum int
	var avg float64
	for {
		var input int
		if _, err := fmt.Scanf("%d", &input); err == io.EOF {
			break
		}

		// 正数不包含 0
		if input < 0 {
			num++
		}
		if input > 0 {
			num2++
			sum += input
		}
	}

	if num2 > 0 {
		//先 float再除，保证精度
		avg = float64(sum) / float64(num2)
	}
	fmt.Printf("%d\n%.1f\n", num, avg)
}
