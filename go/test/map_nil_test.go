package main

import (
	"fmt"
	"testing"
)

var (
	eh *Event_Header
)

type Event_Header struct {
	Extend map[string]string
	Slic   []string
}

func (m *Event_Header) GetExtend() map[string]string {
	if m != nil {
		return m.Extend
	}
	return nil
}

func (m *Event_Header) GetSli() []string {
	if m != nil {
		return m.Slic
	}
	return nil
}

func TestMapNil(t *testing.T) {
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

func TestSliceNil(t *testing.T) {
	if eh == nil {
		t.Log("eh is nil")
		t.Logf("exist val:%s\n", eh.GetSli()[0])
	} else {
		t.Log("eh not nil")
	}
}
