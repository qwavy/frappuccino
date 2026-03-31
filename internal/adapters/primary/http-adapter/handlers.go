package http_adapter

import (
	"frappucchino/internal/services"
)

type Handlers struct {
	Inventory *InventoryHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Inventory: NewInventoryHandler(services.Inventory),
	}
}
