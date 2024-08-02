package main

import (
	"fmt"
	"time"
)

/*
@Time : 2024/7/22 21:31
@Author : echo
@File : chanstruct
@Software: GoLand
@Description:
*/
func worker(done chan struct{}) {
	fmt.Println("working")
	time.Sleep(2 * time.Second)

	close(done)
}
func main() {
	//
	//done := make(chan struct{})
	//go worker(done)
	//<-done
	//fmt.Println("Done")
	a := make(chan int)
	a <- 1 //无gourtine读取会造成死锁
	fmt.Println("1111")
	fmt.Println(<-a)
	fmt.Println("2222")
	close(a)
}
