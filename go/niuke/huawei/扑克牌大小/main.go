package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var(
	ref = map[string]int{
		"A":14,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"10":10,"J":11,"Q":12,"K":13,
	}
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	for{
		bytes, _,err:=reader.ReadLine()
		if err!=nil{
			return
		}

		fmt.Println(comparePoker(string(bytes)))
	}
}

func comparePoker(str string)string{
	arr:=strings.Split(str, "-")
	type1, pokerMap1, lastNum1:=pokerType(strings.Split(arr[0]," "))
	type2, pokerMap2, lastNum2:=pokerType(strings.Split(arr[1]," "))

	if type1 ==6 || type2==6 || type1==4 || type2==4{
		if type1 == 6{
			return arr[0]
		}else if type2 == 6{
			return arr[1]
		}else if type1==4 && type2==4{
			if getNFromMap(pokerMap1, 4) > getNFromMap(pokerMap2,4){
				return arr[0]
			}
			return arr[1]
		}else if type1 == 4{
			return arr[0]
		}else {
			return arr[1]
		}
	}else if type1 != type2{
		return "ERROR"
	}else{
		if type1 == 5{
			if lastNum1> lastNum2{
				return arr[0]
			}
			return arr[1]
		}else if type1 == 3{
			if getNFromMap(pokerMap1,3) > getNFromMap(pokerMap2,3){
				return arr[0]
			}
			return arr[1]
		}else if type1 == 2{
			if getNFromMap(pokerMap1,2) > getNFromMap(pokerMap2,2){
				return arr[0]
			}
			return arr[1]
		}else{
			if getNFromMap(pokerMap1,1) > getNFromMap(pokerMap2,1){
				return arr[0]
			}
			return arr[1]
		}
	}
}

func pokerType(arr []string)(int,map[string]int, int){
	joker:=0
	var lineNum, lastNum int
	count:=map[string]int{}
	for _,n:=range arr{
		count[n]++
		
		if n=="joker" || n== "JOKER"{
			joker++
			if joker==2{
				return 6, nil, 0	//joker
			}
		}else{
			v:=ref[n]
			if lastNum == 0{
				lineNum =1
			} else{
				if lastNum+1==v&&v!=0{
					lineNum++
				}else{
					lineNum=1
				}
			}
			lastNum = v
		}
	}
	if lineNum>=5{
		return 5, count, lastNum
	}

	if getNFromMap(count, 4) != 0{
		return 4, count, 0
	}

	if getNFromMap(count, 3) != 0{
		return 3, count, 0
	}

	if getNFromMap(count, 2) != 0{
		return 2, count, 0
	}

	return 1, count, 0
}

func getNFromMap(m map[string]int, n int)int{
	num:=0
	for k,v := range m{
		if v==n && ref[k]>num{
			num = ref[k]
		}
	}
	return num
}