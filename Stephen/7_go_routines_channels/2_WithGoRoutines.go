package main

import (
	"fmt"
	"net/http"
)

// Writing go in front of function tells go to run it in a separate go routine.
// These routines created by us are called child go routines. They don't get the same amount of respect as the main go routine. This is evident from the fact that if we don't put logic for receiving the message from child go routine, the program doesn't wait for them to complete and simply exit.
// The main routine is the single routine that controls when our program exits or quits

// Q - Do the child go routine stop or continue in the background ?
// They stop when the main goroutine exits. In Go, when the main goroutine (which is the main function) terminates, all other goroutines are also terminated. They do not continue running in the background.

func checkUrlWithRoutine(url string) (e error) {
	resp, err := http.Get((url))
	// Now above call is blocking call. But this time current go routine will tell the other routines - 'Hey I have a blocking call. U can continue if u wanna. I'll let you know when I am done'
	if err != nil {
		fmt.Println(url, " is not up")
		return err
	}
	defer resp.Body.Close()
	fmt.Println(url, " is up")
	return nil
}

func withGoRoutine(urls []string) {
	for _, url := range urls {
		go checkUrlWithRoutine(url)
	}
}
