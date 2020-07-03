package main

/**
 * 准确的算法表述应该是这样的：
设某个真分数的分子为a，分母为b;
把c=(b/a+1)作为分解式中第一个***分数的分母；
将a-b%a作为新的a；
将b*c作为新的b；
如果a等于1，则最后一个***分数为1/b,算法结束；
如果a大于1但是a能整除b，则最后一个***分数为1/(b/a),算法结束；
否则重复上面的步骤。
*/
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string

	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}

		arr := strings.Split(str, "/")
		if len(arr) != 2 {
			return
		}

		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		fmt.Println(calculate(a, b))
	}

}

func calculate(a, b int) string {
	str := ""
	for {
		if a == 1 {
			str = appendNew(str, newFenShu(a, b))
			return str
		}

		if b%(a-1) == 0 {
			str = appendNew(str, newFenShu(1, b/(a-1)))
			a = 1
		} else {
			c := b/a + 1
			a = a - b%a
			b = b * c
			if b%a == 0 {
				b = b / a
				a = 1
			}
			str = appendNew(str, newFenShu(1, c))
		}
	}
}

func appendNew(str string, f fenshu) string {
	if len(str) != 0 {
		str += "+"
	}
	str += fmt.Sprintf("%d/%d", f.f, f.m)
	return str
}

type fenshu struct {
	f, m int
}

func newFenShu(x, y int) fenshu {
	if y%x == 0 && y != 0 {
		y = y / x
		x = 1
	}

	return fenshu{
		f: x,
		m: y,
	}
}

func (f *fenshu) compare(b fenshu) int {
	x := f.f * b.m
	y := b.f * f.m

	if x > y {
		return 1
	} else if x == y {
		return 0
	}
	return -1
}

func (f *fenshu) sub(b fenshu) fenshu {
	x := f.f * b.m
	y := b.f * f.m

	return fenshu{
		f: x - y,
		m: f.m * b.m,
	}
}

func (f *fenshu) isStop() bool {
	return f.f == 1
}
