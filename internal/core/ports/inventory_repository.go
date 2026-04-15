package ports

import (
	"context"
	"frappucchino/internal/core/domain/model/inventory"
)

type InventoryRepository interface {
	GetById(ctx context.Context, id string) (*inventory.Item, error)
	Create(ctx context.Context, inventoryItem *inventory.Item) error
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, id string, inventoryItem *inventory.Item) error
}
