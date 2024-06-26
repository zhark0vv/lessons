package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//syncWaitGroupExample()
	syncAtomicExample()
	//syncOnceExample()
	//syncMutexExample()
	//syncRWMutexExample()
}

func syncWaitGroupExample() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
			fmt.Printf("Goroutine %d finished\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

func syncAtomicExample() {
	var counter int64

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 3)
		}()
	}

	wg.Wait()
	fmt.Printf("Final atomic counter value: %d\n", counter)
}

func syncOnceExample() {
	var once sync.Once

	printOnce := func() {
		fmt.Println("This will be printed only once")
	}

	for i := 0; i < 5; i++ {
		go func() {
			once.Do(printOnce)
		}()
	}
}

func syncMutexExample() {
	var mu sync.Mutex
	var counter int

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Final mutex counter value: %d\n", counter)
}

func syncRWMutexExample() {
	var rwMu sync.RWMutex
	var counter int

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rwMu.RLock()
			fmt.Printf("Goroutine %d reads counter: %d\n", i, counter)
			rwMu.RUnlock()

			rwMu.Lock()
			counter++
			rwMu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final RWMutex counter value: %d\n", counter)
}
