package main

import (
	"fmt"
	"sort"
)

func main() {
	for {
		var n, b int
		num, err := fmt.Scan(&n)
		if num == 0 || err != nil {
			return
		}
		arr := make(ItemSlice, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &arr[i])
		}

		fmt.Scan(&b)
		sort.Sort(arr)
		arr.Print(b == 0)
	}
}

type ItemSlice []int

func (it ItemSlice) Len() int {
	return len(it)
}

func (it ItemSlice) Less(i, j int) bool {
	return it[i] < it[j]
}

func (it ItemSlice) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func (it ItemSlice) Print(asc bool) {
	if len(it) == 0 {
		fmt.Println()
	} else {
		if asc {
			fmt.Printf("%d", it[0])
			for i := 1; i < len(it); i++ {
				fmt.Printf(" %d", it[i])
			}
			fmt.Println()
		} else {
			fmt.Printf("%d", it[len(it)-1])
			for i := len(it) - 2; i >= 0; i-- {
				fmt.Printf(" %d", it[i])
			}
			fmt.Println()
		}
	}
}
