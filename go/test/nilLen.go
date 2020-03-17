package main

import "fmt"

func MapLen(m map[string]struct{}) int {
	return len(m)
}

func SliceLen(s []string) int {
	return len(s)
}

func TestMapLen() {
	var mp map[string]struct{}
	if mp == nil {
		fmt.Printf("nil map len: %d\n", MapLen(mp))
	} else {
		fmt.Printf("map len: %d\n", MapLen(mp))
	}
}

func TestSliceLen() {
	var s []string
	if s == nil {
		fmt.Printf("nil slice len: %d\n", SliceLen(s))
	} else {
		fmt.Printf("slice len: %d\n", SliceLen(s))
	}
}
