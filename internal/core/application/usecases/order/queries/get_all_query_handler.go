package queries

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetAllQueryHandler interface {
	Handle(ctx context.Context) (GetAllResponse, error)
}

type getAllQueryHandler struct {
	db *pgxpool.Pool
}

func NewGetAllQueryHandler(db *pgxpool.Pool) (GetByIdQueryHandler, error) {
	if db == nil {
		return nil, errors.New("db is not provided")
	}

	return &getAllQueryHandler{db: db}, nil
}

func (q *getAllQueryHandler) Handle(ctx context.Context) (GetAllResponse, error) {
	rows, err := q.db.Query(ctx, `SELECT id, customerName, items, status, createdAt FROM order`)
	if err != nil {
		return GetAllResponse{}, err
	}

	var orders []GetAllItem

	for rows.Next() {
		var order GetAllItem

		err = rows.Scan(&order.Id, &order.CustomerName, &order.Items, &order.Status, &order.CreatedAt)
		if err != nil {
			return GetAllResponse{}, err
		}

		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return GetAllResponse{}, err
	}

	response := GetAllResponse{
		orders: orders,
	}

	return response, err
}
