package main

// The sync package provides the WaitGroup struct. It has three primary methods:

// Add(delta int): Increments the counter by delta. Typically, you'd use Add(1) to indicate a new goroutine starting.
// Done(): Decrements the counter by 1. This is called by a goroutine when it finishes its work.
// Wait(): Blocks until the counter reaches 0. This is usually called in the main goroutine to wait for all other goroutines to finish.

// Yes, you must pass a WaitGroup as a reference (pointer) to goroutines. Reason :

// Channels: When you pass a channel by value, a copy of the channel is created. However, this copy refers to the same underlying channel buffer. This allows multiple goroutines to communicate through the same channel.
// WaitGroups: It maintains a counter internally. If you pass a WaitGroup by value to a goroutine, each goroutine would get a copy of the WaitGroup with its own counter. This defeats the purpose of synchronization, as each goroutine would be working on a separate counter.

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the function finishes

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func useWaitGrp() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished")
}
