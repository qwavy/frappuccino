package commands

import (
	"context"
	"errors"
	"frappucchino/internal/core/ports"
)

type DeleteByIdCommandHandler interface {
}

type deleteByIdCommandHandler struct {
	menuRepository ports.MenuRepository
}

func NewDeleteByIdCommandHandler(menuRepository ports.MenuRepository) (DeleteByIdCommandHandler, error) {
	if menuRepository == nil {
		return deleteByIdCommandHandler{}, errors.New("db is not provided")
	}

	return &deleteByIdCommandHandler{menuRepository: menuRepository}, nil
}

func (ch *deleteByIdCommandHandler) Handle(ctx context.Context, command DeleteByIdCommand) error {
	err := ch.menuRepository.DeleteById(ctx, command.Id())

	if err != nil {
		return err
	}
}
