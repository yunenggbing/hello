package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConnect(c net.Conn) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
func main() {

}
