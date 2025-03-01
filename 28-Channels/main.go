package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	
	// Receive Only Channel
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	
	// Send Only Channel
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 0
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}