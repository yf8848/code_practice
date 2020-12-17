package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		bytes, _, err := r.ReadLine()
		if err != nil {
			break
		}

		strs := process(string(bytes))
		for i, s := range strs {
			if i == 0 {
				fmt.Printf("%s", s)
			} else {
				fmt.Printf(" %s", s)
			}
		}
		fmt.Println()
	}
	return
}

func process(str string) []string {
	arr := strings.Split(str, " ")
	res := []string{}

	for i := 1; i < len(arr); i++ {
		s := arr[i]

		for len(s) >= 8 {
			res = append(res, s[:8])
			s = s[8:]
		}
		if len(s) > 0 {
			for len(s) != 8 {
				s += "0"
			}
			res = append(res, s)
		}
	}

	sort.Strings(res)
	return res
}
