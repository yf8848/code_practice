package main

import "fmt"

//fmt.Printf("outer: %d", testElem)
func constIner(){
	const testElem = 10
	fmt.Printf("iner: %d", testElem)
}

func main(){
	constIner()
	//fmt.Printf("outer: %d", testElem)
}
