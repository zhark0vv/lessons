package main

import (
	"fmt"
	"sync"
)

func increment(counter *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		*counter++
	}
}

func main() {
	var (
		counter int
		wg      sync.WaitGroup
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}

	wg.Wait()

	fmt.Println(counter)
}
