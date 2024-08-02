package main

import (
	"fmt"
	"time"
)

/*
@Time : 2024/7/5 17:05
@Author : echo
@File : goRoutineTest
@Software: GoLand
@Description:
*/
func pingPong(in chan int, out chan int) {
	for {
		count := <-in
		out <- count + 1
	}
}
func main() {
	ping := make(chan int)
	pong := make(chan int)
	go pingPong(ping, pong)
	go pingPong(pong, ping)
	ping <- 0
	time.Sleep(1 * time.Second)
	counter := <-ping
	fmt.Println("counter:", counter)

}
