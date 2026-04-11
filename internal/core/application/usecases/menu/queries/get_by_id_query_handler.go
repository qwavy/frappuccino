package queries

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetByIdQueryHandler interface {
}

type getByIdQueryHandler struct {
	db *pgxpool.Pool
}

func NewGetByIdQueryHandler(db *pgxpool.Pool) (GetByIdQueryHandler, error) {
	if db == nil {
		return nil, errors.New("db is not provided")
	}

	return &getByIdQueryHandler{db: db}, nil
}
