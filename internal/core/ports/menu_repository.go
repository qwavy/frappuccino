package ports

import (
	"context"
	"frappucchino/internal/core/domain/model/menu"
)

type MenuRepository interface {
	GetById(ctx context.Context, id string) (*menu.Item, error)
	Create(ctx context.Context, menuItem *menu.Item) error
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, id string, newMenuItem *menu.Item) error
}
