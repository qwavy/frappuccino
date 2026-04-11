package commands

import "errors"

type UpdateByIdCommand struct {
	id           string
	customerName string
	items        []Item
}

func NewUpdateByIdCommand(id string, customerName string) (UpdateByIdCommand, error) {
	if id == "" {
		return UpdateByIdCommand{}, errors.New("id is not provided")
	}

	if customerName == "" {
		return UpdateByIdCommand{}, errors.New("customer name is not provided")
	}

	return UpdateByIdCommand{
		id:           id,
		customerName: customerName,
	}, nil
}

func (c UpdateByIdCommand) Id() string {
	return c.id
}

func (c UpdateByIdCommand) CustomerName() string {
	return c.customerName
}
