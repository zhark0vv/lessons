package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID
	CustomerID uuid.UUID
	Items      []OrderItem
	Status     OrderStatus
	CreatedAt  time.Time
}

// Агрегат - это группа связанных объектов, которые рассматриваются как единое целое

type OrderItem struct {
	ProductID uuid.UUID
	Quantity  int
	Price     float64
}

type OrderStatus string

const (
	Pending   OrderStatus = "PENDING"
	Shipped   OrderStatus = "SHIPPED"
	Delivered OrderStatus = "DELIVERED"
	Cancelled OrderStatus = "CANCELLED"
)

func NewOrder(customerID uuid.UUID, items []OrderItem) *Order {
	return &Order{
		ID:         uuid.New(),
		CustomerID: customerID,
		Items:      items,
		Status:     Pending,
		CreatedAt:  time.Now(),
	}
}

// Событие домена - это событие, которое произошло в предметной области

type OrderCreated struct {
	OrderID    uuid.UUID
	CustomerID uuid.UUID
	Items      []OrderItem
	CreatedAt  time.Time
}

func NewOrderCreated(order *Order) OrderCreated {
	return OrderCreated{
		OrderID:    order.ID,
		CustomerID: order.CustomerID,
		Items:      order.Items,
		CreatedAt:  order.CreatedAt,
	}
}
