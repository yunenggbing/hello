package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
@Time : 2024/7/22 14:19
@Author : echo
@File : clock2
@Software: GoLand
@Description:
*/
func main() {
	port := flag.String("port", "8000", "Port to listen on")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		go handleConn2(conn)
	}

}
func handleConn2(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := fmt.Fprintf(conn, "%s\n", time.Now().Format("15:04:05"))
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}
