package main

import "fmt"

type MapTest struct {
	FilterCountMap map[int]int32 //map[groupID]FilterCount
}

func main() {
	var mt MapTest
	for i := 0; i < 10; i++ {
		mt.FilterCountMap[i]++
	}

	for k, v := range mt.FilterCountMap {
		fmt.Println("%d- %d ", k, v)
	}

}
