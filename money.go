package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
@Time : 2024/7/13 11:57
@Author : echo
@File : money
@Software: GoLand
@Description:
*/
const debug = false

func main() {

	//totalAmount := 20000
	//numPeople := 70
	//amounts := make([]int, numPeople)
	//for i := range amounts {
	//	amounts[i] = 1
	//}
	//totalAmount -= numPeople
	//rand.Seed(time.Now().UnixNano())
	//for totalAmount > 0 {
	//	index := rand.Intn(numPeople)
	//	amounts[index]++
	//	totalAmount--
	//}
	//for i, amount := range amounts {
	//	fmt.Printf("Person %d: %.2f元\n", i+1, float64(amount)/100)
	//}

	var w io.Writer
	fmt.Printf("(%T, %[1]v)\n", w)
	w = os.Stdout
	w = io.Writer(os.Stdout)
	fmt.Printf("(%T, %[1]v)\n", w)
	w.Write([]byte("hello, writer\n")) // 向屏幕输出了hello, writer
	fmt.Printf("(%T, %[1]v)\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("(%T, %[1]v)\n", w)
	w.Write([]byte("hello")) // 向缓冲区写入了hello
	fmt.Printf("(%T, %[1]v)\n", w)

	w = nil
	fmt.Printf("(%T, %[1]v)\n", w)
}

func f1(out io.Writer) {
	if out != nil {
		out.Write([]byte("done！\n"))
	}
}
