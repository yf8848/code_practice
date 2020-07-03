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

		s, d := process(str)
		fmt.Printf("%s,%d\n", s, d)
	}
}

func process(str string) (string, int) {
	var longStr, tmpStr string
	var long, tmp int

	for _, c := range str {
		if unicode.IsNumber(c) {
			tmpStr += string(c)
			tmp++
		} else {
			if long < tmp {
				longStr = tmpStr
				long = tmp
			} else if long == tmp {
				longStr += tmpStr
			}
			tmpStr = ""
			tmp = 0
		}
	}

	if long < tmp {
		longStr = tmpStr
		long = tmp
	} else if long == tmp {
		longStr += tmpStr
	}

	return longStr, long
}
