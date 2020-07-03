package main

import "fmt"

func main(){
	var n int
	for{
		num,err:=fmt.Scan(&n)
		if num==0 || err!= nil{
			return
		}
		fmt.Printf("%d\n", int(2*n+n*(n-1)*3/2))
	}
}