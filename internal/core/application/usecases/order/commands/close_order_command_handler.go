package commands

import (
	"context"
	"errors"
	"frappucchino/internal/core/ports"
)

type CloseOrderCommandHandler interface {
}

type closeOrderCommandHandler struct {
	orderRepository ports.OrderRepository
}

func NewCloseOrderCommandHandler(orderRepository ports.OrderRepository) (CloseOrderCommandHandler, error) {
	if orderRepository == nil {
		return nil, errors.New("order repository is not provided")
	}

	return &closeOrderCommandHandler{orderRepository: orderRepository}, nil
}

func (ch *closeOrderCommandHandler) Handle(ctx context.Context, command CloseOrderCommand) error {
	order, err := ch.orderRepository.GetById(ctx, command.Id())

	if err != nil {
		return err
	}

	return order.Close()
}
