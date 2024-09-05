package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(baseCtx, 2*time.Second)
	defer cancel()

	ctxWithValue := context.WithValue(ctxWithTimeout, "key", "value")

	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Горутина завершена")
		case <-ctxWithValue.Done():
			fmt.Println("Контекст отменен:", ctxWithValue.Err(),
				ctxWithValue.Value("key"))

		}
	}()

	time.Sleep(5 * time.Second)
}
