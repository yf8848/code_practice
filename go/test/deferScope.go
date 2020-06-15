package main

import "fmt"

func TestDerferScope() {
	defer fmt.Println("defer scope test finish")
	fmt.Sprintln("defer scope test 1 start.")

	{
		defer fmt.Println("defer scope test 1")
	}

	fmt.Sprintln("defer scope test 1 start.")

	{
		defer fmt.Println("defer scope test 2")
	}
}
