package main

import (
	"fmt"
	"syscall"
)

func errSyscall() error {
	err := syscall.Chmod(":invalid path:", 0666)
	if err != nil {
		fmt.Printf("err:%s, errno:%d\n", err, err.(syscall.Errno))
		return err
	}
	return nil
}

func main() {
	errSyscall()
}
