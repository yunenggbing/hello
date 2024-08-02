package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	fmt.Println(buf)

}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done\n"))
	}
}
