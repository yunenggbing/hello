package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("value1:", i)
		}()
	}
	wg.Wait()
	fmt.Println("----------")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println("value:", j)
		}(i)
	}
	wg.Wait()
	fmt.Println("-----Over-----")
	//time.Sleep(2 * time.Second)
	// 使用 fmt.Println 输出到标准输出
	//fmt.Println("All Arguments:", os.Args)
	//now := time.Now()
	//
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}
	//
	//end := time.Now()
	//fmt.Println("所用时间；", "now:", "end:", end.Sub(now), now, end)
	//fmt.Println(s)
	//now2 := time.Now()
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//end2 := time.Now()
	//fmt.Println("所用时间；", "now:", "end:", end2.Sub(now2), now2, end2)
}
