package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2024/7/31 14:36
@Author : echo
@File : onceDemo
@Software: GoLand
@Description:
*/

var once = sync.Once{}

func Init() {
	fmt.Println("init")
}

func main() {
	for i := 0; i < 100; i++ {
		go once.Do(Init)
	}

}
