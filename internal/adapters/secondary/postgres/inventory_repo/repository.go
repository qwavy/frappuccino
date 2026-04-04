package inventory_repo

import (
	"context"
	"frappucchino/internal/core/domain/model/inventory"
	"frappucchino/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type inventoryRepository struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) ports.InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetAll(ctx context.Context) ([]*inventory.Item, error) {
	rows, err := r.db.Query(ctx, `SELECT ingredient_id, name, quantity, unit, price, created, updated FROM inventory`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*inventory.Item

	for rows.Next() {
		var itemDTO InventoryItemDTO
		err = rows.Scan(&itemDTO.IngredientID, &itemDTO.Name, &itemDTO.Quantity, &itemDTO.Unit, &itemDTO.Price, &itemDTO.Created, &itemDTO.Updated)

		if err != nil {
			return nil, err
		}

		items = append(items, DTOtoDomain(itemDTO))
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *inventoryRepository) GetById(ctx context.Context, id string) (*inventory.Item, error) {
	rows := r.db.QueryRow(ctx, `SELECT ingredient_id, name, quantity, unit, price, created, updated FROM inventory WHERE ingredient_id = $1`, id)

	var itemDTO InventoryItemDTO

	err := rows.Scan(&itemDTO.IngredientID, &itemDTO.Name, &itemDTO.Quantity, &itemDTO.Unit, &itemDTO.Price, &itemDTO.Created, &itemDTO.Updated)

	if err != nil {
		return nil, err
	}

	return DTOtoDomain(itemDTO), nil
}

func (r *inventoryRepository) Create(ctx context.Context, inventoryItem *inventory.Item) error {
	inventoryItemDTO := DomainToDTO(inventoryItem)
	_, err := r.db.Exec(ctx, `INSERT INTO inventory (ingredient_id, name, quantity, unit, price, created, updated) values ($1, $2, $3, $4,$5,$6,$7)`, inventoryItemDTO.IngredientID, inventoryItemDTO.Name, inventoryItemDTO.Quantity, inventoryItemDTO.Unit, inventoryItemDTO.Price, inventoryItemDTO.Created, inventoryItemDTO.Updated)
	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepository) DeleteById(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM inventory WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepository) UpdateById(ctx context.Context, id string, inventoryItem *inventory.Item) error {
	inventoryItemDTO := DomainToDTO(inventoryItem)
	_, err := r.db.Exec(ctx, `UPDATE inventory SET name = $2,  price = $3, unit = $4, quantity = $5, updated_at = $6`, id, inventoryItemDTO.Name, inventoryItemDTO.Price, inventoryItemDTO.Unit, inventoryItemDTO.Quantity, inventoryItemDTO.Updated)
	if err != nil {
		return err
	}
	return nil
}
