package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := r.ReadLine()
		if len(bytes) == 0 {
			break
		}

		processStr(bytes)
	}
}

func processStr(bytes []byte) {
	arr := make([]int, 256)
	for _, c := range bytes {
		if unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) || unicode.IsSpace(rune(c)) {
			arr[c]++
		}
	}

	for {
		max := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] == 0 {
				continue
			}
			if arr[i] > arr[max] {
				max = i
			}
		}
		if arr[max] == 0 {
			break
		}
		fmt.Printf("%c", max)
		arr[max] = 0
	}
	fmt.Println()
}
