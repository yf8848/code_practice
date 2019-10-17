package main

import (
	"fmt"
)

func noErr() error {
	return nil
}

func byErr() error {
	return fmt.Errorf("error")
}

func main() {
	err := noErr()
	if err == nil {
		err = byErr()
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}
