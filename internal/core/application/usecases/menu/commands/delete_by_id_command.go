package commands

import "errors"

type DeleteByIdCommand struct {
	id string
}

func NewDeleteByIdCommand(id string) (DeleteByIdCommand, error) {
	if id == "" {
		return DeleteByIdCommand{}, errors.New("id is not provided")
	}

	return DeleteByIdCommand{
		id: id,
	}, nil
}

func (c DeleteByIdCommand) Id() string {
	return c.id
}
