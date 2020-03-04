package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	arg := os.Args[1]
	f, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, f)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
