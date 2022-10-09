package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Read write mutex in GO")

	wg := &sync.WaitGroup{}
	mux := &sync.RWMutex{}
	// iifs or immediately invoked functions. Are functions which can be defined and executed at the same time.
	// Syntax func(){}()
	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mux)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mux)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mux)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		// Apply a read lock where ever you are reading from a common memory
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		wg.Done()
	}(wg, mux)

	wg.Wait()
	fmt.Println(score)
}
