package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, troila.")
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			io.WriteString(resp, "hello, troila.")
		} else {
			io.WriteString(resp, "405")
		}
	})
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println("start failed")
		os.Exit(1)
	}
}
