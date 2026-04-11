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
	orderRepository ports.OrderRepository
}

func NewDeleteByIdCommandHandler(orderRepository ports.OrderRepository) (DeleteByIdCommandHandler, error) {
	if orderRepository == nil {
		return nil, errors.New("order repository is not provided")
	}

	return &deleteByIdCommandHandler{orderRepository: orderRepository}, nil
}

func (ch *deleteByIdCommandHandler) Handle(ctx context.Context, command DeleteByIdCommand) error {
	order, err := ch.orderRepository.GetById(ctx, command.Id())

	if err != nil {
		return err
	}

	if order.Status() == "closed" {
		return errors.New("cant delete closed order")
	}

	err = ch.orderRepository.DeleteById(ctx, command.Id())

	if err != nil {
		return err
	}

	return err
}
