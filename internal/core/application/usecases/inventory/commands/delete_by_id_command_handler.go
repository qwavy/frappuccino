package commands

import (
	"context"
	"errors"
	"frappucchino/internal/core/ports"
)

type DeleteByIdCommandHandler interface {
	Handle(ctx context.Context, command DeleteByIdCommand) error
}
type deleteByIdCommandHandler struct {
	inventoryRepository ports.InventoryRepository
}

func NewDeleteByIdCommandHandler(inventoryRepository ports.InventoryRepository) (DeleteByIdCommandHandler, error) {
	if inventoryRepository == nil {
		return nil, errors.New("inventory repository is not provided")
	}

	return &deleteByIdCommandHandler{
		inventoryRepository: inventoryRepository,
	}, nil
}
func (ch *deleteByIdCommandHandler) Handle(ctx context.Context, command DeleteByIdCommand) error {
	inventoryItem, err := ch.inventoryRepository.GetById(ctx, command.Id())

	inventoryItem.
}
