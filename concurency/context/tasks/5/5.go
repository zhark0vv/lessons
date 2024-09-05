// Создать свободную реализацию errorgroup
// *, если останется время: добавить метод, создающий группу, мерджущую ошибки в канал

// Важные требования:
// - При первой ошибке все горутины должны быть отменены
// - Ошибка должна быть возвращена из Wait
// - Сстояние контекста должно быть сохранено при передаче в горутины

package main

import (
	"context"
	"sync"
)

// ErrorGroup позволяет управлять группой горутин и собирать ошибки.
type ErrorGroup struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	mu     sync.Mutex
	err    error
}

// NewErrorGroup создает новый ErrorGroup с заданным контекстом.
func NewErrorGroup(ctx context.Context) *ErrorGroup {
	childCtx, cancel := context.WithCancel(ctx)
	return &ErrorGroup{
		ctx:    childCtx,
		cancel: cancel,
	}
}

// Go запускает новую горутину и добавляет ее в группу.
func (g *ErrorGroup) Go(f func(ctx context.Context) error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(g.ctx); err != nil {
			g.mu.Lock()
			defer g.mu.Unlock()
			if g.err == nil {
				g.err = err
				g.cancel() // Отменить все горутины при первой ошибке
			}
		}
	}()
}

// Wait ожидает завершения всех горутин и возвращает первую ошибку.
func (g *ErrorGroup) Wait() error {
	g.wg.Wait()
	return g.err
}
