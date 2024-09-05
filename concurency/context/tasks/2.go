/*
Есть функция processDataInternal, которая может выполняться неопределенно долго.
Чтобы контролировать процесс, мы добавили таймаут выполнения через context. Какие недостатки кода ниже?
*/

package main

import (
	"context"
	"io"
)

type Service struct{}

func (s *Service) ProcessData(timeoutCtx context.Context, r io.Reader) error {
	errCh := make(chan error)

	go func() {
		errCh <- s.processDataInternal(r)
	}()

	select {
	case err := <-errCh:
		return err
	case <-timeoutCtx.Done():
		return timeoutCtx.Err()
	}
}

func (s *Service) processDataInternal(r io.Reader) error {
	// Simulate work
	_, err := r.Read([]byte{0})
	return err
}
