package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	m := map[string][]int{}
	less := map[string]int{}
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		str := string(bytes)
		pre := str[0 : len(str)-4]
		last := str[len(str)-3:]

		//fmt.Println(pre, last)
		if _, ok := m[pre]; !ok {
			m[pre] = []int{}
		}

		t := string2int(last)
		m[pre] = append(m[pre], t)

		if v, ok := less[pre]; !ok {
			less[pre] = t
		} else if t < v {
			less[pre] = t
		}
	}
	res := 0
	for k, v := range m {
		res += count(v, less[k])
	}

	fmt.Println(res)
}

func count(v []int, k int) int {
	res := 0
	for i := 0; i < len(v); i++ {
		if v[i] == k {
			res++
		}
	}
	return res
}

func string2int(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		t := int(str[i] - '0')
		res = res*10 + t
	}
	return res
}
