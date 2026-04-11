package queries

import "errors"

type GetByIdQuery struct {
	id string
}

func NewGetByIdQuery(id string) (GetByIdQuery, error) {
	if id == "" {
		return GetByIdQuery{}, errors.New("id is required")
	}

	return GetByIdQuery{}, nil
}
