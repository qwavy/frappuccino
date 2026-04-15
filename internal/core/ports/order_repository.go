package ports

import (
	"context"
	"frappucchino/internal/core/domain/model/order"
)

type OrderRepository interface {
	GetById(ctx context.Context, id string) (*order.Order, error)
	Create(ctx context.Context, menuItem *order.Order) error
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, id string, newOrder *order.Order) error
}
