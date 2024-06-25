package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
