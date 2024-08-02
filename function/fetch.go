package function

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

func fetch1(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, err
}
func f(x int) {
	fmt.Printf("f(%d)\n),x+0/x")
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])

}
