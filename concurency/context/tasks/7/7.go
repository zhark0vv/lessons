/*
Задача: Написать функцию на Go, которая принимает на вход слайс целых чисел и размер батча,
и возвращает итоговый результат — сумму всех чисел, обработанных в батчах параллельно.
*/

package main

import (
	"fmt"
	"sync"
)

func processBatch(batch []int, wg *sync.WaitGroup, resultCh chan<- int) {
	defer wg.Done()

	sum := 0
	for _, num := range batch {
		sum += num
	}

	resultCh <- sum
}

func parallelBatchProcessing(data []int, batchSize int) int {
	var wg sync.WaitGroup
	resultCh := make(chan int, len(data)/batchSize+1)

	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]
		wg.Add(1)
		go processBatch(batch, &wg, resultCh)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	totalSum := 0
	for sum := range resultCh {
		totalSum += sum
	}

	return totalSum
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	batchSize := 5
	result := parallelBatchProcessing(data, batchSize)
	fmt.Println("Total sum:", result)
}
