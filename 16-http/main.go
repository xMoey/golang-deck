package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://golang.org/pkg/net/http/#Response
var resp http.Response

func main() {
	// https://golang.org/pkg/net/http/
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// remember: an array of bytes is a string
	// to manually init an array, read the resp into it and print:
	// mk := make([]byte, 99999)
	// resp.Body.Read(mk)
	// fmt.Println(string(mk))

	// https://golang.org/pkg/io/#Writer
	io.Copy(os.Stdout, resp.Body)
}