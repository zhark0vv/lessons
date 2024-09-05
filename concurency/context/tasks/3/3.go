package main

import (
	"context"
	"fmt"
	"time"
)

func proc1() {
	timer := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ctx, cancel2 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel2()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Работа завершена")
	case <-ctx.Done():
		fmt.Println("Контекст отменен:", ctx.Err(), time.Since(timer))
	}
}

func proc2() {
	timer := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx, cancel2 := context.WithTimeout(ctx, 1*time.Second)
	defer cancel2()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Работа завершена")
	case <-ctx.Done():
		fmt.Println("Контекст отменен:", ctx.Err(), time.Since(timer))
	}
}

func main() {
	proc1()
	proc2()
}
