package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2024/7/5 15:34
@Author : echo
@File : onceDemo
@Software: GoLand
@Description: sync.Once：1、线程安全 2、高性能 3、更简洁；需要确保 Do 方法中的函数不会因为错误而提前退出
*/
type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Printf("Creating singleton instance now.")
		instance = &singleton{}
	})
	return instance
}
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ins := GetInstance()
			fmt.Printf("Instance address: %p\n", ins)
		}()
	}
	wg.Wait()
}
