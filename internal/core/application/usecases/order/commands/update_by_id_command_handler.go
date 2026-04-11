package commands

import (
	"context"
	"errors"
	"frappucchino/internal/core/ports"
)

type UpdateByIdCommandHandler interface {
}

type updateByIdCommandHandler struct {
	orderRepository ports.OrderRepository
}

func NewUpdateByIdCommandHandler(orderRepository ports.OrderRepository) (UpdateByIdCommandHandler, error) {
	if orderRepository == nil {
		return nil, errors.New("order repository is not provided")
	}

	return &updateByIdCommandHandler{orderRepository: orderRepository}, nil
}

func (ch *updateByIdCommandHandler) Handle(ctx context.Context, command UpdateByIdCommand) error {
	order, err := ch.orderRepository.GetById(ctx, command.id)

	if err != nil {
		return err
	}

	err = order.SetCustomerName(command.customerName)

	if err != nil {
		return err
	}

}
