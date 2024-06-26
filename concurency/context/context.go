package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	contextBackgroundExample()
	contextTODOExample()
	contextWithCancelExample()
	contextWithDeadlineExample()
	contextWithTimeoutExample()
}

func contextBackgroundExample() {
	ctx := context.Background()
	doWork(ctx, "Background")
}

func contextTODOExample() {
	ctx := context.TODO()
	doWork(ctx, "TODO")
}

func contextWithCancelExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		doWork(ctx, "WithCancel")
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func contextWithDeadlineExample() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	doWork(ctx, "WithDeadline")
}

func contextWithTimeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	doWork(ctx, "WithTimeout")
}

func doWork(ctx context.Context, name string) {
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("%s: working...\n", name)
		case <-ctx.Done():
			fmt.Printf("%s: done, reason: %v\n", name, ctx.Err())
			return
		}
	}
}
