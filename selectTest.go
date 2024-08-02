package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

/*
@Time : 2024/7/2 15:19
@Author : echo
@File : selectTest
@Software: GoLand
@Description: select 练习
*/
type Html interface {
	io.ReadCloser
	Getdoc(url string) (err error)
	ForeachNode(n *html.Node, pre, post func(n *html.Node))
	SearchNode(tag string) []*html.Node
	PrintAll()
}

// 不对外暴露，一切操作均通过方法实现
// 除了Html接口中定义的方法, 还可以自己拓展方法, 例如Show方法
type htmlParser struct {
	Html
	str  string
	node *html.Node
	resp *http.Response
}

// 自定义读入
func (doc *htmlParser) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, fmt.Errorf("error,read data is empty")
	}
	n = copy(p, doc.str)
	if doc.node != nil {
		return 0, fmt.Errorf("error,node is not empty")
	}
	doc.node, err = html.Parse(strings.NewReader(doc.str))
	return n, err
}

// clear 清空
func (doc *htmlParser) Clear() {
	doc.node = nil
	doc.resp = nil
}

// close 关闭连接
func (doc *htmlParser) Close() error {
	return doc.resp.Body.Close()
}

// Getdoc 通过url获取网页内容
func (doc *htmlParser) Getdoc(url string) (err error) {
	doc.resp, err = http.Get(url)
	defer doc.Close()
	if doc.resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http get error:%s", doc.resp.Status)
		return err
	}
	doc.node, err = html.Parse(doc.resp.Body)
	if err != nil {
		return fmt.Errorf("http parse error:%s", err)
	}
	return nil
}

// ForEachNode 遍历节点
func (doc *htmlParser) ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if n == nil {
		return
	}
	if pre != nil {
		pre(n)
	}
	doc.ForEachNode(n.FirstChild, pre, post)
	doc.ForEachNode(n.NextSibling, pre, post)
	if post != nil {
		post(n)
	}
}

// SearchNode 搜索节点
func (doc *htmlParser) SearchNode(tag string) []*html.Node {
	var nodes []*html.Node
	doc.ForEachNode(doc.node, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			nodes = append(nodes, n)
		}
	}, nil)
	return nodes
}

func (doc *htmlParser) PrintAll() {
	doc.ForEachNode(doc.node, func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%s\n", n.Data)
		}
	}, nil)
}
func (doc *htmlParser) Show() {
	fmt.Println(doc.resp)
	fmt.Println()
	fmt.Println(doc.node)
}
func main() {
	var doc htmlParser
	doc.Getdoc("http://www.baidu.com")
	doc.Show()
}
