package queries

import (
	"context"
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

func (q *getByIdQueryHandler) Handle(ctx context.Context, query GetByIdQuery) (GetByIdResponse, error) {
	row := q.db.QueryRow(ctx, `SELECT id, customerName, items, status, createdAt FROM order WHERE id = $1`, query.id)

	var orderDTO GetByIdResponse
	err := row.Scan(&orderDTO.Id, &orderDTO.CustomerName, &orderDTO.Items, &orderDTO.Status, &orderDTO.CreatedAt)
	if err != nil {
		return GetByIdResponse{}, err
	}

	return orderDTO, nil
}
