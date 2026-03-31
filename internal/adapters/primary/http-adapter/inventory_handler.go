package http_adapter

import (
	"frappucchino/internal/services"
	http2 "frappucchino/pkg/http"
	"net/http"
)

type InventoryHandler struct {
	inventoryService services.InventoryService
}

func NewInventoryHandler(inventoryService services.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
	}
}

func (h *InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	inventoryItems, err := h.inventoryService.GetAll()
	if err != nil {
		http2.SendHttpError(w, err)
	}

	http2.WriteJSON(w, 200, inventoryItems)
}

func (h *InventoryHandler) GetById(w http.ResponseWriter, r *http.Request) {
	inventoryItem, err := h.inventoryService.GetById()
}
