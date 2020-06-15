package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Printf("%s\n", revStr(str))
}

func revStr(str string) string {
	new := ""
	// 跳过行尾的‘\n’
	for i := len(str) - 2; i >= 0; i-- {
		new += string(str[i])
	}
	return new
}
