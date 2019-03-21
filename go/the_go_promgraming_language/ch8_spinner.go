package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

/*
	func for outer main func test fib
*/
func TestMain() {
	go spinner(1000 * time.Microsecond)
	fmt.Println("\rfilN: ", fib(45))
}
