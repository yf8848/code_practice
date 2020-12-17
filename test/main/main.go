package main

import "fmt"

func main() {

	for i := range get() {
		fmt.Printf("%d\n", i)
	}
	fmt.Println("finish")
}

func get() []int {
	return nil
}
