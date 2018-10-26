package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":9090", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello web!")
	}))
}
