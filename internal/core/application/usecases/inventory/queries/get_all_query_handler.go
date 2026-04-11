package queries

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetAllQueryHandler interface {
	Handle(ctx context.Context) (GetAllQueryResponse, error)
}
type getAllQueryHandler struct {
	db *pgxpool.Pool
}

func NewGetAllQueryHandler(db *pgxpool.Pool) (GetAllQueryHandler, error) {
	if db == nil {
		return nil, errors.New("db is not provided")
	}
	return &getAllQueryHandler{db: db}, nil
}

func (q *getAllQueryHandler) Handle(ctx context.Context) (GetAllQueryResponse, error) {
	rows, err := q.db.Query(ctx, `SELECT ingredient_id, name, quantity, unit, price, created, updated FROM inventory`)
	if err != nil {
		return GetAllQueryResponse{}, err
	}
	defer rows.Close()

	var inventoryItems []Item

	for rows.Next() {
		var item Item
		err = rows.Scan(&item.IngredientID, &item.Name, &item.Quantity, &item.Unit, &item.Price, &item.Created, &item.Updated)

		if err != nil {
			return GetAllQueryResponse{}, err
		}

		inventoryItems = append(inventoryItems, item)
	}

	err = rows.Err()

	if err != nil {
		return GetAllQueryResponse{}, err
	}
	response := GetAllQueryResponse{
		InventoryItems: inventoryItems,
	}

	return response, nil
}
