package main

import (
	"fmt"
	"hello/link"
	"log"
	"os"
)

/*
@Time : 2024/7/1 16:39
@Author : echo
@File : craw
@Software: GoLand
@Description:
*/

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := link.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

/*
*
优化
*/
var tokens = make(chan struct{}, 20)

func craw2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := link.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//	func main() {
//		fmt.Println(os.Args[1:])
//		workList := make(chan []string)
//		var n int
//		n++
//		go func() {
//			workList <- os.Args[1:]
//		}()
//		seen := make(map[string]bool)
//		for ; n > 0; n-- {
//			list := <-workList
//			for _, l := range list {
//				if !seen[l] {
//					seen[l] = true
//					go func(link string) {
//						workList <- craw2(link)
//					}(l)
//				}
//			}
//		}
//	}
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := craw2(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
