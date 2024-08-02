package main

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
@Time : 2024/7/25 16:16
@Author : echo
@File : firstHttp
@Software: GoLand
@Description:
*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.PostForm("/user", url.Values{"name": {"echo"}, "age": {"18"}})
	http.ListenAndServe(":8080", nil)
}
