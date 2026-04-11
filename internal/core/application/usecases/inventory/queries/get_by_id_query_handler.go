package queries

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetByIdQueryHandler interface {
	Handle(ctx context.Context, query GetByIdQuery) (GetByIdResponse, error)
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
	rows := q.db.QueryRow(ctx, `SELECT ingredient_id, name, quantity, unit, price, created, updated FROM inventory WHERE ingredient_id = $1`, query.Id())

	var inventoryItem GetByIdResponse

	err := rows.Scan(&inventoryItem.IngredientID, &inventoryItem.Name, &inventoryItem.Quantity, &inventoryItem.Unit, &inventoryItem.Price, &inventoryItem.Created, &inventoryItem.Updated)

	if err != nil {
		return GetByIdResponse{}, err
	}

	response := GetByIdResponse{
		IngredientID: inventoryItem.IngredientID,
		Name:         inventoryItem.Name,
		Quantity:     inventoryItem.Quantity,
		Unit:         inventoryItem.Unit,
		Price:        inventoryItem.Price,
		Created:      inventoryItem.Created,
		Updated:      inventoryItem.Updated,
	}

	return response, nil
}
