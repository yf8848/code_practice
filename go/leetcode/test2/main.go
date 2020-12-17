package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		bytes, _, err := r.ReadLine()
		if err != nil {
			break
		}

		fmt.Printf("%s\n", process(string(bytes)))
	}
}

func process(str string) string {
	need := map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	win := make(map[byte]int, 0)

	right := 0
	start, end := 0, 0
	max := 0
	match := 0

	for right < len(str) {
		pre := right
		if right > 0 {
			pre = right - 1
		}
		t := str[pre]
		c := str[right]
		right++

		if need[c] || (isAlpha(c) && !isAlpha(t) && win[c] == 0) {
			win[c]++
			match++
		} else {
			//fmt.Printf("%s max:%d match:%d\n", str[left:right-1], max, match)
			if max <= match && match > 0 {
				start = right - match - 1
				end = right - 1
				//left = right - 1
				max = match
				match = 0
			} else {
				//left = right - 1
			}
		}
	}
	if max <= match && match > 0 {
		start = right - match - 1
		end = right - 1
		//left = right - 1
		max = match
		match = 0
	}
	return str[start:end]
}

func isAlpha(c byte) bool {
	return c == '+' || c == '-' || c == '.'
}
