package postgres

import (
	"context"
	"frappucchino/internal/core/domain/model/inventory"
	"frappucchino/internal/core/ports"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type inventoryRepository struct {
	db *pgxpool.Pool
}

type InventoryDTO struct {
	ingredientID int
	name         string
	quantity     float64
	unit         Unit
	price        float64
	created      time.Time
	updated      time.Time
}

func NewInventoryRepository(db *pgxpool.Pool) ports.InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetAll(ctx context.Context) ([]inventory.Item, error) {
	rows, err := r.db.Query(ctx, "SELECT ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan()

		if err != nil {
			return err
		}
	}
	return nil, nil
}

func (r *inventoryRepository) GetById(ctx context.Context) ([]inventory.Item, error) {
	return nil, nil
}

func (r *inventoryRepository) Create(ctx context.Context) ([]inventory.Item, error) {
	return nil, nil

}

func (r *inventoryRepository) DeleteById(ctx context.Context) ([]inventory.Item, error) {
	return nil, nil

}

func (r *inventoryRepository) UpdateById(ctx context.Context, id string, newInventoryItem inventory.Item) error {
	//	getById

	//inventoryItem, err := r.db.QueryRow()

	return nil
}
