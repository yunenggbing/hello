package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@Time : 2024/7/5 15:07
@Author : echo
@File : mutexDemo
@Software: GoLand
@Description:
*/
func main() {
	var mu sync.Mutex
	var count int
	var wg sync.WaitGroup
	//启动100个 groutine 来增加计数
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()   //获取锁
			count++     //只有一个groutine。可以执行这行代码
			mu.Unlock() //释放锁
		}()
	}
	wg.Wait() //等待所有groutine完成
	fmt.Println("count:", count)
	var x, y int
	go func() {
		x = 1
		fmt.Println("y:", y)
	}()
	go func() {
		y = 1
		fmt.Println("x:", x)
	}()
	time.Sleep(time.Second)
}
