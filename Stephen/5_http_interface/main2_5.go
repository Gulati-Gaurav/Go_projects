package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type PrintConsole struct {
}

func (PrintConsole) Write(b []byte) (n int, err error) {
	fmt.Println(string(b))
	return len(b), nil
}

func AnotherWay1_5() {
	// Let's create our own Write implementing type

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Found error", err)
		os.Exit(1)
	}

	pc := PrintConsole{}
	io.Copy(pc, resp.Body)
}
