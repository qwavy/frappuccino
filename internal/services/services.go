package services

import (
	"frappucchino/internal/adapters/secondary/postgres"
)

type InventoryService interface {
}

type Services struct {
	Inventory InventoryService
}

func NewServices(repositories *postgres.Repositories) *Services {
	return &Services{Inventory: NewInventoryService(repositories.InventoryRepository)}
}
