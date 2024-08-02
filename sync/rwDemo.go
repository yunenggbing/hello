package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2024/7/31 16:14
@Author : echo
@File : rwDEmo
@Software: GoLand
@Description:
*/

var (
	value int
	rwMu  sync.RWMutex
)

func readValue(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.RLock()
	defer rwMu.RUnlock()
	fmt.Println("Read Value:", value)
}
func writeValue(wg *sync.WaitGroup, newValue int) {
	defer wg.Done()
	rwMu.Lock()
	defer rwMu.Unlock()
	value = newValue
	fmt.Println("Write Value:", newValue)
}

func main() {
	var wg sync.WaitGroup
	value = 0
	wg.Add(3)

	go readValue(&wg)
	go writeValue(&wg, 1)
	go readValue(&wg)
	wg.Wait()

}
