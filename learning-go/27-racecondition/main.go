package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in GO")

	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	// iifs or immediately invoked functions. Are functions which can be defined and executed at the same time.
	// Syntax func(){}()
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mux)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mux)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mux)

	wg.Wait()
	fmt.Println(score)
}
