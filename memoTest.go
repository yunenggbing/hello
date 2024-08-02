package main

import (
	"fmt"
	"hello/memo"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
@Time : 2024/7/5 15:59
@Author : echo
@File : memoTest
@Software: GoLand
@Description:
*/
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func main() {
	m := memo.New(httpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()

}
