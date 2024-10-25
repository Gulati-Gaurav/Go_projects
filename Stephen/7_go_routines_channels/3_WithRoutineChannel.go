package main

import (
	"fmt"
	"net/http"
)

// solution - Channels or WaitGroups
// Channel are used to communicate between multiple running routines.
// we will use channel to make sure main routine knows when the child have completed there work and only then exit the program
// Think of channel as a whatsapp group. Whenever a routine sends a message to the channel, all routines that have access to the channel receive the message (not only main).
// Channels are typed. Means can only share one type of information amoung routines.

// Q - if one channel consumes a message then can other channel also do so ?
// No, in Go, once a message is consumed from a channel by one goroutine, it is no longer available for other goroutines to consume. A message in a channel can only be received once.

// Buffered vs Unbuffered channels
// Unbuffered channels: Require a corresponding receiver for every send operation. If there's no receiver ready, the sender blocks.
// Buffered channels: Can store a limited number of values before blocking. This allows senders to continue without waiting for immediate reception.

// In Go, a goroutine attempting to send a value to an unbuffered channel will block until another goroutine is ready to receive it.
// Unbuffered channels: If no receiver is ready, the sender blocks.
// Buffered channels: Can store a limited number of values before blocking.
// If all goroutines are blocked trying to send to a channel and no goroutine is ready to receive, a deadlock occurs.

// ch2 := make(chan int, 2) // to make channel with buffer in cases when you don't want number of listener = number of speaker

func checkUrlWithRoutineChannel(url string, c chan string) (e error) {
	resp, err := http.Get((url))
	if err != nil {
		fmt.Println(url, " is not up")
		c <- "Might be down I think"
		return err
	}
	defer resp.Body.Close()
	fmt.Println(url, " is up")

	// This is how to send message to channel
	c <- "Yup it's up"
	// Now go tells other routines that we have received a message over the channel, is anyone waiting for it ?

	return nil
}

func withGoRoutineChannel(urls []string) {
	c := make(chan string)
	// This is how we create a channel.
	// Treat it just as any other value inside our application. Same scoping rules.
	// Now we have to send the channel to the function of routine so that function/routine has access tp channel
	for _, url := range urls {
		go checkUrlWithRoutineChannel(url, c)
	}

	// This is how we consume messages from channels.
	// Wait for a value to be sent to channel. When we get one, assign it to myString
	// myString := <- c
	// Or if you wanna wait for a value to be sent to the channel and when you get one you log that out immediately.
	// fmt.Println(<-c)

	for i := range urls {
		fmt.Println(<-c, i) // to wait for all child routine. This is a blocking call.
		// all blocking calls make that routine to sleep
	}

	// Or
	// familar loop
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c, i)
	// }
}

// Test of whether main keeps spinning new or lets the responded routine to complete - I created 1000 links and it seems it spinned new routines first of all and then let the routine be completed. Either routine creation is too fast or it doesn't give time to other. Can't add more links as rate limiter at these websites.
