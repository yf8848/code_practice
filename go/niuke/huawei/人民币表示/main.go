package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ling = 0
	shi  = 10
	bai  = 10 * shi
	qian = 10 * bai
	wan  = 10 * qian
	yi   = wan * wan

	numRef = map[int]string{
		0: "零", 1: "壹", 2: "贰", 3: "叁", 4: "肆", 5: "伍", 6: "陆", 7: "柒", 8: "捌", 9: "玖",
		shi: "拾", bai: "佰", qian: "仟", wan: "万", yi: "亿",
	}
)

func main() {
	var num float64
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		str := fmt.Sprintf("%.2f", num)
		strArr := strings.Split(str, ".")
		zhen, err := strconv.Atoi(strArr[0])
		if err != nil {
			return
		}
		dot, err := strconv.Atoi(strArr[1])
		if err != nil {
			return
		}
		fmt.Printf("人民币%s\n", process(zhen, dot))
	}
}

func process(zhen, dot int) string {
	str := processZhen(zhen)
	find := fmt.Sprintf("%s%s", numRef[1], numRef[shi])
	if strings.HasPrefix(str, find) {
		str = strings.Replace(str, find, numRef[shi], 1)
	}
	if len(str) > 0 {
		str += "元"
	}

	dotStr := processDot(dot)
	if len(dotStr) == 0 && len(str) > 0 {
		str += "整"
	}
	str += dotStr
	return str
}

func processDot(dot int) string {
	str := ""
	if dot > 0 {
		jiao := dot / 10
		fen := dot % 10

		if jiao > 0 {
			str += numRef[jiao] + "角"
		}

		if fen > 0 {
			str += numRef[fen] + "分"
		}
	}
	return str
}

func processZhen(zhen int) string {
	str := ""
	if zhen > yi {
		num := zhen / yi
		str += processZhen(num)
		str += numRef[yi]
		zhen = zhen % yi
	}

	if zhen > wan {
		num := zhen / wan
		str += processZhen(num)
		str += numRef[wan]
		zhen = zhen % wan
	}

	if zhen > qian {
		num := zhen / qian
		str += processZhen(num)
		str += numRef[qian]
		zhen = zhen % qian
	}
	if zhen > bai {
		num := zhen / bai
		str += processZhen(num)
		str += numRef[bai]
		zhen = zhen % bai
	}

	if zhen > shi {
		num := zhen / shi
		str += processZhen(num)
		str += numRef[shi]
		zhen = zhen % shi
	}

	if zhen > ling {
		str += numRef[zhen]
	}

	return str
}
