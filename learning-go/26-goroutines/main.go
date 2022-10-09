package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// Wait group for waiting goroutine thread to finish
var wg sync.WaitGroup

var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://lco.dev",
		"https://google.com",
		"https://bing.com",
		"https://yandex.com",
		"https://youtube.com",
	}

	for _, web := range websitelist {
		// go keyword is used to create a thread or goroutine
		// goroutines are handled by go runtime.
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	defer wg.Done()
	if err != nil {
		fmt.Println("Incorrect endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
