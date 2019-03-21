package main

import (
	"io"
	"log"
	"net"
	"time"
)

/*
test server func
*/
func TestClockServer() {

	listener, err := net.Listen("tcp", "127.0.0.1:8005")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		//go handleConClock(conn)
		go HandleEcho(conn)
	}
}

func handleConClock(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))

		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
