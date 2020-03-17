package main

import (
	"fmt"
	"runtime"
	"time"
)

func handlerRecover(err interface{}) {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]

	if err != nil {
		text := fmt.Sprintf("panic stack: %s", buf)
		fmt.Printf("handler panic recover: %s\n", text)
		time.Sleep(time.Second)
		panic(err)
	}
}

func goWithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			//handlerRecover(err)
		}
		handlerRecover(nil)
	}()

	fn()
}

func TestPanic() {
	go goWithRecover(func() { panic("panic") })
	time.Sleep(1 * time.Second)
	fmt.Println("I stand.")
}
