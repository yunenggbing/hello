package main

import (
	"log"
	"net"
	"os"
)

/*
@Time : 2024/7/1 11:19
@Author : echo
@File : netcat2
@Software: GoLand
@Description:
*/
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
