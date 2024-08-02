package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// reader 结构体用于限制 io.Reader 的读取字节数。
type reader struct {
	i    int       // 已经读取的字节数。
	read io.Reader // 原始的 io.Reader 对象。
	n    int       // 允许读取的总字节数。
}

// Read 方法实现了 io.Reader 接口，用于限制读取的字节数。
func (r *reader) Read(p []byte) (n int, err error) {
	// 如果已经读取的字节数达到允许的最大值，返回 EOF 错误。
	if r.i >= r.n {
		return 0, io.EOF
	}
	// 如果剩余可读取的字节数少于请求的字节数，只读取剩余的字节数，并返回 EOF 错误。
	if r.n-r.i < len(p) {
		n, _ = r.read.Read(p[:r.n-r.i])
		err = io.EOF
		return
	}
	// 正常读取数据，更新已读取的字节数。
	n, err = r.read.Read(p)
	if err != nil {
		return
	}
	r.i += n
	return
}

// LimitReader 函数用于创建一个限制读取字节数的 io.Reader 接口实现。
// 它接受一个原始的 io.Reader 对象和一个限制的字节数，返回一个新的 io.Reader 对象。
func LimitReader(r io.Reader, n int) io.Reader {
	return &reader{i: 0, read: r, n: n}
}

func main() {
	// 创建一个字符串阅读器，并限制读取的字节数。
	r := strings.NewReader("123456789")
	r1 := LimitReader(r, 4)
	// 使用 bufio.Scanner 从限制后的阅读器中逐行读取。
	s := bufio.NewScanner(r1)
	// 设置 Scanner 的分割方式为按字节分割。
	s.Split(bufio.ScanBytes)
	// 遍历读取的每一行，并打印。  s.Scan 会使用到io.Read方法，所以此处会调用上方的Read方法
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
