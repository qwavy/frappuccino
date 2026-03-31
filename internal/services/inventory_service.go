package services

import (
	"errors"
	"frappucchino/internal/adapters/secondary/postgres"
	"frappucchino/internal/core/domain/model/inventory"
	"frappucchino/internal/core/ports"
	"log/slog"
)

type inventoryService struct {
	inventoryRepo ports.InventoryRepository
}

func NewInventoryService(inventoryRepo ports.InventoryRepository) InventoryService {
	return &inventoryService{inventoryRepo: inventoryRepo}
}

func (s *inventoryService) GetAll() (inventory.Item, error) {
	return nil, nil
}

func (s *inventoryService) GetById(id string) (*inventory.InventoryItem, error) {
	inventoryItem, err := s.InventoryRepository.GetById(id)
	if err != nil {
		slog.Error("service: failed to get inventory item", "id", id, "error", err)
		return nil, err
	}

	return inventoryItem, nil
}

func (s *InventoryService) Create(item inventory.InventoryItem) error {
	if err := item.Validate(); err != nil {
		slog.Warn("service: validation failed for inventory creation", "id", item.IngredientID, "error", err)
		return err
	}

	_, err := s.InventoryRepository.GetById(item.IngredientID)

	if err == nil {
		slog.Warn("inventory item already exists", "id", item.IngredientID)
		return inventory.InventoryItemAlreadyExists
	}

	if !errors.Is(err, inventory.InventoryItemNotFound) {
		slog.Error("service: failed to get inventory item", "id", item.IngredientID, "error", err)
		return err
	}

	if err := s.InventoryRepository.Create(item); err != nil {
		slog.Error("service: repository failed to create item", "id", item.IngredientID, "error", err)
		return err
	}

	slog.Info("service: inventory item created successfully", "id", item.IngredientID)
	return nil
}

func (s *InventoryService) Update(id string, item inventory.InventoryItem) error {
	if err := item.Validate(); err != nil {
		slog.Warn("service: validation failed for inventory update", "id", id, "error", err)
		return err
	}

	if err := s.InventoryRepository.UpdateById(id, item); err != nil {
		slog.Error("service: repository failed to update item", "id", id, "error", err)
		return err
	}

	slog.Info("service: inventory item updated successfully", "id", id)
	return nil
}

func (s *InventoryService) DeleteById(id string) error {
	if err := s.InventoryRepository.DeleteById(id); err != nil {
		slog.Error("service: failed to delete inventory item", "id", id, "error", err)
		return err
	}

	slog.Info("service: inventory item deleted successfully", "id", id)
	return nil
}
