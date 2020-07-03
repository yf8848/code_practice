package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}

		fmt.Printf("%s\n", process(str))
	}
}

func process(str string) string {
	newStr := ""
	numStr := ""
	for _, c := range str {
		if unicode.IsNumber(c) {
			numStr += string(c)
		} else {
			if len(numStr) > 0 {
				newStr += "*" + numStr + "*"
				numStr = ""
			}
			newStr += string(c)
		}
	}
	if len(numStr) > 0 {
		newStr += "*" + numStr + "*"
	}
	return newStr
}
