package ports

import (
	"context"
	"frappucchino/internal/core/domain/model/inventory"
)

type InventoryRepository interface {
	GetAll(ctx context.Context) ([]inventory.Item, error)
}
