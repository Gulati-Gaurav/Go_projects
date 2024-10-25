package main

import (
	"fmt"
	"net/http"
)

// When our program starts, it runs on a single main go routine. Think of it as a thread. Or a engine that executes code.
// That routine keeps executing whatever code it sees line by line. When a blocking call like the http.Get comes, it continues to wait for it at that line blocking the further code execution.
// In short words, processes get executed in a serial order, one after other.

// This is a problem since we will have to wait a lot in case of large number of websites.
// But in our case we can run multiple checkUrl in parallel since they don't rely on each other

func checkUrl(url string) (e error) {
	resp, err := http.Get(url) // this is a blocking call as the main go routine has to wait for it.
	// http.Get is a synchronous call

	if err != nil {
		fmt.Println(url, " is not up")
		return err
	}
	defer resp.Body.Close()
	fmt.Println(url, " is up")
	return nil
}

func withoutRoutines(urlList []string) {
	for _, url := range urlList {
		checkUrl(url)
	}
}
