package main

import "fmt"

func main() {
	str := ""

	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s -> %d", str, longestValidParentheses(str))
	}
	return
}
