package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/*
@Time : 2024/7/2 15:43
@Author : echo
@File : du
@Software: GoLand
@Description: 并发的目录遍历
*/
func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}

}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

/*
*
输出路径硬盘空间
*/
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func main() {
	//main1()
	//main2()
	main3()
}

/*
*
方式1
*/
func main1() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	filseSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, filseSize)
		}
		close(filseSize)
	}()
	var nfiles, nbytes int64
	for size := range filseSize {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main2() {
	now := time.Now()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	filseSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, filseSize)
		}
		close(filseSize)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-filseSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
	fmt.Println("end----", time.Since(now))
}

func main3() {
	now := time.Now()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	filseSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, filseSize)
	}
	go func() {
		n.Wait()
		close(filseSize)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-filseSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
	fmt.Println("end----", time.Since(now))
}
