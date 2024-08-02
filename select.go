package main

import (
	"fmt"
	"time"
)

/*
@Time : 2024/7/22 16:01
@Author : echo
@File : select
@Software: GoLand
@Description:
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan int)
	go func() {
		ch1 <- 1
		done <- 1
	}()
	go func() {
		ch2 <- 2
		done <- 2
	}()
	for {
		select {
		case <-ch1:
			fmt.Println("ch1")
		case <-ch2:
			fmt.Println("ch2")
		case v := <-done:
			fmt.Println("done", v)
			if v == 2 {
				fmt.Println("value is 2 exit")
				return
			}
		default:
			fmt.Println("default")
			time.Sleep(1 * time.Second)
		}
	}

}
