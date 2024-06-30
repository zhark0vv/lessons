package main

import (
	"sync"
)

func getTasks() []string {
	// Запрос к API - движку задач
	return make([]string, 0)
}

func processTask(wg *sync.WaitGroup, task string) {
	defer wg.Done()
	// Задача отправляет что-то в API
}

func main() {
	// Задачи получаются из внешнего API
	tasks := getTasks()

	wg := sync.WaitGroup{}
	for _, t := range tasks {
		wg.Add(1)
		go processTask(&wg, t)
	}

	wg.Wait()
}
