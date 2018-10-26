package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.ListenAndServe(":9090", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello web!")
		//which one is the fastest
		bw := bufio.NewWriter(w)
		bw.WriteString("hello web!\n")
		bw.Flush()
		//which one is the fastest
		io.Copy(w, strings.NewReader("hello web!\n"))
	}))
}
