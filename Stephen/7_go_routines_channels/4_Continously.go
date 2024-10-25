package main

import (
	"fmt"
	"net/http"
	"time"
)

// Now send link into the channel

func checkUrlContinously(url string, c chan string) (e error) {
	resp, err := http.Get((url))
	if err != nil {
		fmt.Println(url, " is not up")
		c <- url
		defer resp.Body.Close()
		return err
	}
	defer resp.Body.Close()
	fmt.Println(url, " is up")

	c <- url

	return nil
}

func withContinously(urls []string) {
	c := make(chan string)

	for _, url := range urls {
		go checkUrlContinously(url, c)
	}

	// syntax to run the loop till eternity
	// for {
	// 	go checkUrlContinously(<-c, c)
	// }

	// Or another way to make the code more readable
	for l := range c { // here we are saying wait for the channel to return a value and then assign it to l. Hence a blocking call

		// It is also important to put the sleep in the function call and not in the main go routine else would mean that main go routine would be making 1 go routine every 5 seconds.

		go func(link string) { // imp NEVER EVER in practice we should never refer to same variable in >=2 routines
			time.Sleep(time.Second * 3) // sleeping the current thread for 3 seconds
			go checkUrlContinously(link, c)
		}(l) // this is function literal. Like the lambda function in c#
	}
}

// Is there any issue with the following code?

// package main

// func main() {
//      c := make(chan string)
//      c <- "Hi there!"
// }

// The syntax of this program is OK, but the program will go into a deadlock because non one is listening to the message we are passing into channel
