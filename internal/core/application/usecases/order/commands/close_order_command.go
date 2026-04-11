package commands

import "errors"

type CloseOrderCommand struct {
	id string
}

func NewCloseOrderCommand(id string) (CloseOrderCommand, error) {
	if id == "" {
		return CloseOrderCommand{}, errors.New("id is not provided")
	}

	return CloseOrderCommand{
		id: id,
	}, nil
}

func (c CloseOrderCommand) Id() string {
	return c.id
}
