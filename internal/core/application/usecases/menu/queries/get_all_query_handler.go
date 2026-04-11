package queries

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetAllQueryHandler interface {
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

func (q *getAllQueryHandler) Handle(ctx context.Context, query GetAllQuery) (GetAllResponse, error) {
	rows, err := q.db.Query(ctx, `SELECT id, name, description, price, itemSize, allergens, categories, customization, ingredients from menu`)

	if err != nil {
		return GetAllResponse{}, err
	}
	defer rows.Close()

	var menuItems GetAllResponse

	for rows.Next() {
		var menuItemDTO Item
		err = rows.Scan(&menuItemDTO.Id, &menuItemDTO.Name, &menuItemDTO.Description, &menuItemDTO.Price, &menuItemDTO.ItemSize, &menuItemDTO.Allergens, &menuItemDTO.Categories, &menuItemDTO.Customization, &menuItemDTO.Ingredients)
		if err != nil {
			return GetAllResponse{}, err
		}
	}

	err = rows.Err()

	if err != nil {
		return GetAllResponse{}, err
	}

	return menuItems, nil
}
