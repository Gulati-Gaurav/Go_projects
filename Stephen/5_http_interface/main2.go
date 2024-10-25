package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func AnotherWay1() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Found error", err)
		os.Exit(1)
	}

	// Similar to Reader, we have a Writer interface which takes []byte array and converts it to desired output.
	// os.Stdout is of Writer type.
	// See Images
	byteSlice := make([]byte, 99999)
	n, err2 := resp.Body.Read(byteSlice)
	if err2 != nil {
		fmt.Println(n, err2) // same EOF error
	}
	os.Stdout.Write(byteSlice)

	// or shortest way. No use of byte slice
	// Copy helps in transfer
	io.Copy(os.Stdout, resp.Body)
}
