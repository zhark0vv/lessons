package service

import (
	"context"
	"errors"
)

//go:generate mockery --name DataProvider --output ./mocks --filename data_provider.go
type DataProvider interface {
	GetData(ctx context.Context, id int) (string, error)
}

type Service struct {
	provider DataProvider
}

func New(provider DataProvider) *Service {
	return &Service{provider: provider}
}

func (s *Service) ProcessData(ctx context.Context, id int) (string, error) {
	data, err := s.provider.GetData(ctx, id)
	if err != nil {
		return "", err
	}

	if data == "" {
		return "", errors.New("data not found")
	}

	return "Processed: " + data, nil
}
