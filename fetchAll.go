package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start  to goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%0.2fs  elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send  to channel ch
		return
	}
	fileName, _ := strings.CutPrefix(url, "https://")
	out, err := os.Create(fileName + ".txt")
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(out, resp.Body)
	err = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%0.2fs  %7d  %s", secs, nbytes, url)
}
