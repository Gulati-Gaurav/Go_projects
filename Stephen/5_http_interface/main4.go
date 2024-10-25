package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func AnotherWay3() {

	// Accept-encoding can be
	// gzip
	// compress
	// deflate
	// br
	// identity
	// *

	client := &http.Client{}
	req, err2 := http.NewRequest(
		"GET",
		"http://google.com",
		http.NoBody,
	)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	req.Header.Set("Accept-Encoding", "identity") // NOTE THIS LINE
	res, err3 := client.Do(req)
	if err3 != nil {
		fmt.Println(err3)
		os.Exit(1)
	}
	body, err4 := io.ReadAll(res.Body)
	res.Body.Close()
	if err4 != nil {
		fmt.Println(err4)
		os.Exit(1)
	} else {
		fmt.Println(string(body))
	}
	_ = body

}
