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

	// another way - gives the response in bytes 
	body, err5 := io.ReadAll(resp.Body)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}

// So basically 4 methods
// 1. make byte array yourself
// 2. use the io.Copy() if have a writer ready
// 3. use the io.ReadAll() if wanna play with response yourself.