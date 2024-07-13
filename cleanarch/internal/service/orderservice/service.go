package orderservice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"lessons/cleanarch/internal/domain"
)

// Анти-коррупционный слой для интеграции с репозиторием продуктов

type productRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	Save(ctx context.Context, product *domain.Product) error
}

type eventPublisher interface {
	Publish(ctx context.Context, event interface{}) error
}

type OrderService struct {
	productRepository productRepo
	eventPublisher    eventPublisher
}

func NewOrderService(productRepo productRepo, publisher eventPublisher) *OrderService {
	return &OrderService{
		productRepository: productRepo,
		eventPublisher:    publisher,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, customerID uuid.UUID, items []domain.OrderItem) (*domain.Order, error) {
	for _, item := range items {
		product, err := s.productRepository.FindByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s", product.Name)
		}
		product.Stock -= item.Quantity
		if err := s.productRepository.Save(ctx, product); err != nil {
			return nil, err
		}
	}

	order := domain.NewOrder(customerID, items)

	// Отправка события о создании заказа
	event := domain.NewOrderCreated(order)
	// Отправка события в шину сообщений
	if err := s.eventPublisher.Publish(ctx, event); err != nil {
		return nil, err
	}
	return order, nil
}
