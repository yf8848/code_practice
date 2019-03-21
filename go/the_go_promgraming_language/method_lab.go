package main

import "fmt"

type methodLab struct {
	name string
	age  int
}

var Person methodLab

func (pson *methodLab) tell() *methodLab {
	fmt.Printf("hello. %v\n", pson)
	return pson
}

func (pson *methodLab) addAge() *methodLab {
	pson.age++
	return pson
}

// redeclare
// func (pson *methodLab) tell() {
// 	fmt.Println("hello, this is indictor. %v", pson)
// }

func InitPs() methodLab {
	a := methodLab{name: "non-indictor", age: 1}
	fmt.Printf("Init non-indictor addr: %p\n", a)
	return a
}

func InitPsIndctor() *methodLab {
	a := new(methodLab)
	fmt.Printf("Init indictor addr: %p\n", a)
	return a
}

func TestMethod() {
	no := InitPs()
	no.tell()
	no.addAge()
	no.tell()
	in := InitPsIndctor()
	in.tell()
	in.addAge()
	in.tell()
}
