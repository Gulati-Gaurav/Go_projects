package main

import (
	"fmt"
	"sync"
)

func mutex() {
	mut := &sync.Mutex{}
	rwmut := &sync.RWMutex{}
	fmt.Println(mut)
	fmt.Println(rwmut)

	// mutex is for proper lock
	// rwmutex allows you you define whether read lock or write lock
}
