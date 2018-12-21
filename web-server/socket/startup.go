package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := "127.0.0.1:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
