package function

import (
	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	_, _, _ = fetch1("111")
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
