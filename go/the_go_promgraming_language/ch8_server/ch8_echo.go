package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

/*
* handle conn for echo 3
 */
func HandleEcho(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Print(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Println(c, "\t", shout)
	time.Sleep(delay)
	fmt.Println(c, "\t", strings.ToLower(shout))
}
