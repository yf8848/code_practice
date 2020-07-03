package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main()  {
	reader:=bufio.NewReader(os.Stdin)
	for{
		bytes,_, err:=reader.ReadLine()
		if err!= nil{
			return 
		}

		fmt.Println(checkPasswd(string(bytes)))
	}
}


func checkPasswd(str string)string{
	var num ,lower,upper,other int
	for _,c:=range str{
		if unicode.IsNumber(c){
			num++
		}else if unicode.IsLower(c){
			lower++
		}else if unicode.IsUpper(c){
			upper++
		}else{
			other++
		}
	}

	score:=0
	if len(str)<=4{
		score+=5
	}else if len(str)<=7{
		score+=10
	}else{
		score+=25
	}

	if lower>0 && upper > 0{
		score+=20
	}else if lower>0 || upper>0{
		score+=10
	}

	if num>1{
		score+=20
	}else if num>0{
		score+=10
	}

	if other>1{
		score+=25
	}else if other>0{
		score+=10
	}

	if num>0&& lower>0 &&upper>0&&other>0{
		score+=5
	}else if num>0&& lower+upper>0&&other>0{
		score+=3
	}else if num>0&& lower+upper>0{
		score+=2
	}
	if score >=90{
		return "VERY_SECURE"
	}else if score >=80{
		return "SECURE"
	}else if score >=70{
		return "VERY_STRONG"
	}else if score>=60{
		return "STRONG"
	}else if score>=50{
		return "AVERAGE"
	}else if score >=25{
		return 		"WEAK"
	}
	return "VERY_WEAK"
}