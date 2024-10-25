package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	// WaitGroup()
	// mutex()
	// raceCondition()
	channels()
}

func WaitGroup() {
	wg := &sync.WaitGroup{}
	wg.Add(4) // number of times you wanna wait for signal
	go hitAPI("https://www.youtube.com", wg)
	go hitAPI("https://www.amazon.com", wg)
	go hitAPI("https://www.google.com", wg)
	go hitAPI("https://www.facebook.com", wg)
	wg.Wait() // wait for signals
}

func hitAPI(url string, wg *sync.WaitGroup) {
	defer wg.Done() // signal done

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}
