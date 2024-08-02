package main

import (
	"fmt"
	"os"
	"time"
)

/*
@Time : 2024/7/2 15:08
@Author : echo
@File : countdown
@Software: GoLand
@Description:
*/
func main() {
	fmt.Println("Commencing countdown.  Press return to abort.")
	//tick := time.Tick(1 * time.Second)
	//for countdown := 10; countdown > 0; countdown-- {
	//	fmt.Println(countdown)
	//	<-tick
	//}
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
		launch()
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

}

func launch() {
	fmt.Println("Launching  over...")

}
