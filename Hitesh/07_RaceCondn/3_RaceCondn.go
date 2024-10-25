package main

import (
	"fmt"
	"sync"
)

func raceCondition() {
	// To identify possible threats due to race condition, run the command
	// go run --race .  (include .)

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	rwm := &sync.RWMutex{}
	func() {}() // IFFE

	score := []int{0}

	// No guarantee that the order of the execution will be 1, 2, 3
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex, rwm *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Routine 1")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
	}(wg, m, rwm)
	go func(wg *sync.WaitGroup, m *sync.Mutex, rwm *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Routine 2")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
	}(wg, m, rwm)
	go func(wg *sync.WaitGroup, m *sync.Mutex, rwm *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Routine 3")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
	}(wg, m, rwm)

	wg.Wait()
}
