package main

import (
	"fmt"
	"strconv"
)

func TestString2Num() {
	strNum := "10000"
	NagtiveNum := "-20000"

	num, err := strconv.ParseUint(strNum, 10, 64)
	if err != nil {
		fmt.Printf("ParseUint %s err: %v ", strNum, err)
	}

	nNum, err := strconv.ParseUint(NagtiveNum, 10, 64)
	if err != nil {
		fmt.Printf("ParseUint %s err: %v ", NagtiveNum, err)
	}

	fmt.Printf("%s: %d, %s:%d", strNum, num, NagtiveNum, nNum)
}
