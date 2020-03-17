package main

import (
	"fmt"
	"strconv"
)

func TestList() {
	list := make([]int, 10)

	for i := 0; i < 10; i++ {
		list = append(list, i*10)
	}

	fmt.Println(list)
}

func TestEmptySlice() {
	var rcmirrorTasks []string
	var i int64 = 0
	for i < 10 {
		i++
		str := strconv.FormatInt(i, 10)
		rcmirrorTasks = append(rcmirrorTasks, str+"_test")
	}

	fmt.Printf("list: %v", rcmirrorTasks)
}
