package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workerPoolExample()
	doneChannelExample()
	countingSemaphoreExample()
	fanInFanOutExample()
}

func workerPoolExample() {
	const numWorkers = 3
	tasks := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	taskChan := make(chan int, len(tasks))

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(taskChan, &wg, i)
	}

	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	wg.Wait()
	fmt.Println("Worker pool example completed")
}

func worker(tasks <-chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", workerID, task)
		time.Sleep(time.Second) // Simulate work
	}
}

func doneChannelExample() {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Received done signal")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("Done channel example completed")
}

func countingSemaphoreExample() {
	const maxGoroutines = 2
	semaphore := make(chan struct{}, maxGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			semaphore <- struct{}{}
			fmt.Printf("Processing task %d\n", taskID)
			time.Sleep(time.Second) // Simulate work
			<-semaphore
		}(i)
	}

	wg.Wait()
	fmt.Println("Counting semaphore example completed")
}

func fanInFanOutExample() {
	inputChan := make(chan int)
	outputChan := make(chan int)
	const numWorkers = 3

	var wg sync.WaitGroup

	// Fan-out
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range inputChan {
				fmt.Printf("Worker %d processing task %d\n", workerID, task)
				outputChan <- task * 2
			}
		}(i)
	}

	// Fan-in
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	// Sending tasks
	go func() {
		for i := 1; i <= 5; i++ {
			inputChan <- i
		}
		close(inputChan)
	}()

	// Receiving results
	for result := range outputChan {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println("Fan-in/Fan-out example completed")
}
