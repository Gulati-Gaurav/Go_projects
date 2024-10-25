package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func AnotherWay2() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Found error", err)
		os.Exit(1)
	}

	// another way
	body, err5 := io.ReadAll(resp.Body)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
