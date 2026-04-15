package queries

import "errors"

type GetByIdQuery struct {
	id string
}

func NewGetByIdQuery(id string) (GetByIdQuery, error) {
	if id == "" {
		return GetByIdQuery{}, errors.New("id is not provided")
	}

	return GetByIdQuery{
		id: id,
	}, nil
}

func (q GetByIdQuery) Id() string {
	return q.id
}
