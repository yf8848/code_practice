package main

import "fmt"

var (
	eh *Event_Header
)

type Event_Header struct {
	Extend map[string]string
}

func (m *Event_Header) GetExtend() map[string]string {
	if m != nil {
		return m.Extend
	}
	return nil
}

func main() {
	if eh == nil {
		fmt.Println("eh is nil")
		if val, ok := eh.GetExtend()["ext"]; ok {
			fmt.Printf("exist val:%s\n", val)
		} else {
			fmt.Printf("not exist key:%s， val:%v\n", "ext", val)
		}
	} else {
		fmt.Printf("eh not nil")
	}

	return
}
