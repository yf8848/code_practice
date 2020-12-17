package mian

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		bytes, _, err := r.ReadLine()
		if err != nil {
			break
		}

		fmt.Printf("%d\n", process(string(bytes)))
	}
	return
}

func process(str string) int {
	arr := strings.Split(str, " ")

	nums := make([]int, len(arr))
	for i, s := range arr {
		num, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return 0
		}

		if s[len(s)-1:] == "Y" {
			nums[i] = num
		} else {
			nums[i] = num * 7
		}
	}

	sort.Ints(nums)

	return nums[len(nums)-1] - nums[0]
}
