package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2024/7/30 15:26
@Author : echo
@File : syncDemo.go
@Software: GoLand
@Description:
sync.WaitGroup :  wg.Add(10000)  wg.Done()  wg.wait()
sync.Mutex : 互斥锁  lock()和unLock()
sync.RWMutex : 读写互斥锁  lock()和unLock()  RLock()和RUnLock()
sync.Once :  sync.Once{}  Do()  只执行一次
*/
var sum = 0
var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
var rwMutex = sync.RWMutex{}

func main() {
	wg.Add(10000)
	for i := 1; i <= 10000; i++ {

		go func() {
			defer wg.Done()
			add()
		}()

	}
	//time.Sleep(time.Second * 2)
	wg.Wait() //等待所有携程执行完毕
	fmt.Println(sum)

}

func add() {
	//mutex.Lock()
	sum += 1
	//defer mutex.Unlock()

}
