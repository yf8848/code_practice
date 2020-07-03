package main

import (
	"fmt"
	"strings"
)

func main() {
	var short, long string
	for {
		_, err := fmt.Scan(&short, &long)
		if err != nil {
			return
		}

		if strings.Contains(long, short) {
			fmt.Println("true")
		} else {
			istotal := true
			ref := map[rune]bool{}
			for _, c := range long {
				ref[c] = true
			}

			for _, c := range short {
				if !ref[c] {
					istotal = false
					break
				}
			}

			if istotal {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		}
	}
}
