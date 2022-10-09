package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO")

	myCh := make(chan int, 4)
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(5)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	// Another syntax in channel: chan<- is Only sending channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup, m *sync.Mutex) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg, mut)
	// Only receiving channel <-chan
	// In receiving channels, closing of channel is not allowed
	go func(ch <-chan int, wg *sync.WaitGroup, m *sync.Mutex) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg, mut)
	// Only receiving channel <-chan
	go func(ch <-chan int, wg *sync.WaitGroup, m *sync.Mutex) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg, mut)

	wg.Wait()
}
