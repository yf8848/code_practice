package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}

		fmt.Println(IsIp(str))
	}
}

func IsIp(str string) string {
	arr := strings.Split(str, ".")
	if len(arr) != 4 {
		return "NO"
	}

	for _, c := range arr {
		if n, err := strconv.Atoi(c); err != nil || n < 0 || n >= math.MaxUint8 {
			return "NO"
		}
	}
	return "YES"
}
