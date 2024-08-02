package link

import "golang.org/x/net/html"

/*
@Time : 2024/7/1 16:45
@Author : echo
@File : outline
@Software: GoLand
@Description:
*/
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
