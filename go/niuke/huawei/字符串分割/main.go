package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		num, err := fmt.Scanf("%d", &n)
		if num == 0 || err != nil {
			return
		}
		for i := 0; i < n; i++ {
			var str string
			fmt.Scanf("%s", &str)
			fmtPrint(str)
		}
	}

}

func fmtPrint(str string) {
	for len(str) >= 8 {
		fmt.Printf("%s\n", str[0:8])
		str = str[8:]
	}
	if len(str) > 0 {
		fmt.Printf("%s\n", fill0(str))
	}
}

func fill0(str string) string {
	for len(str) < 8 {
		str += "0"
	}
	return str
}
