package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string
	var k int
	for {
		_, err := fmt.Scan(&str1)
		if err != nil {
			return
		}
		_, err = fmt.Scan(&str2)
		if err != nil {
			return
		}
		_, err = fmt.Scan(&k)
		if err != nil {
			return
		}

		strm := strings.Split(str1, ",")
		strn := strings.Split(str2, ",")
		if len(strm) != len(strn) {
			return
		}

		m := toIntArr(strm)
		n := toIntArr(strn)

		fmt.Println(getMoney(m, n, k))
	}

}

type em struct {
	score int
	idx   int
}

type elems []em

func (s elems) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s elems) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s elems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getMoney(m, n []int, k int) int {

	ref := make(map[int]int, len(m))
	rget := make(map[int]int, len(m))
	get := make(elems, len(m))
	for i := 0; i < len(m); i++ {
		ref[i] = m[i]
		get[i] = em{
			score: n[i] - m[i],
			idx:   i,
		}
		rget[n[i]-m[i]] = i
	}

	sort.Sort(get)

	for {
		lower := math.MaxInt64
		for _, v := range get {
			if m[v.idx] <= k {
				k += n[v.idx] - m[i]
				delete(ref, i)
			} else {
				if m[i] < lower {
					lower = m[i]
				}
			}
		}

		if len(ref) == 0 || lower > k {
			return k
		}
	}
}

func toIntArr(strs []string) []int {
	res := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = string2int(strs[i])
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
