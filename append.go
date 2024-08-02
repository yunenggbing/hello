package main

import (
	"fmt"
)

func main() {
	//var x, y []int
	//for i := 0; i < 10; i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	//	x = y
	//}
	data := []string{"", "1", "2", ""}
	fmt.Println(nonempty(data))
	fmt.Println(data)
	//fmt.Println(appendInt(x, y))

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap > 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	//i := 0
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			//strings[i] = s
			// i++
			out = append(out, s)
		}
	}
	return out
}
