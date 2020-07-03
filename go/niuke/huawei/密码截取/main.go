package main

import "fmt"

func main(){
	var str string
	for {
		n, err:=fmt.Scan(&str)
		if n==0|| err!=nil{
			return 
		}

		fmt.Println(longestPalindrome(str))
	}
}

func longestPalindrome(str string)int{
	maxLen:=0
	for i :=range str{
		if i-maxLen>=1 && str[i-maxLen-1:i+1] == reverse(str[i-maxLen-1:i+1]){
			maxLen +=2
			continue
		}

		if i-maxLen>=0&&str[i-maxLen:i+1] == reverse(str[i-maxLen:i+1]){
			maxLen++
		}
	}
	return maxLen
}

func reverse(s string) string{
	str:=[]byte(s)
	for i,j:=0, len(str)-1; i<j; i,j=i+1, j-1{
		str[i],str[j] = str[j],str[i]
	}
	return string(str)
}