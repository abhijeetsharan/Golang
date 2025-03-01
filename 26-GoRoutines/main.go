package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mutex sync.Mutex // pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://golang.org",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
