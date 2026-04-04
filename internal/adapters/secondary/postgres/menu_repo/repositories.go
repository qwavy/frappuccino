package menu_repo

import (
	"context"
	"frappucchino/internal/core/domain/model/menu"
	"frappucchino/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type menuRepository struct {
	db *pgxpool.Pool
}

func NewMenuRepository(db *pgxpool.Pool) ports.MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) GetAll(ctx context.Context) ([]*menu.Item, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, description, price, itemSize, allergens, categories, customization, ingredients from menu`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuItems []*menu.Item

	for rows.Next() {
		var menuItemDTO MenuItemDTO
		err = rows.Scan(&menuItemDTO.Id, &menuItemDTO.Name, &menuItemDTO.Description, &menuItemDTO.Price, &menuItemDTO.ItemSize, &menuItemDTO.Allergens, &menuItemDTO.Categories, &menuItemDTO.Customization, &menuItemDTO.Ingredients)
		if err != nil {
			return nil, err
		}

		menuItems = append(menuItems, DTOtoDomain(menuItemDTO))
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return menuItems, nil
}
func (r *menuRepository) GetById(ctx context.Context, id string) (*menu.Item, error) {
	rows := r.db.QueryRow(ctx, `SELECT id, name, description, price, itemSize, allergens, categories, customization, ingredients from menu WHERE id = $1`, id)

	var menuItemDTO MenuItemDTO

	err := rows.Scan(&menuItemDTO.Id, &menuItemDTO.Name, &menuItemDTO.Description, &menuItemDTO.Price, &menuItemDTO.ItemSize, &menuItemDTO.Allergens, &menuItemDTO.Categories, &menuItemDTO.Customization, &menuItemDTO.Ingredients)
	if err != nil {
		return nil, err
	}

	return DTOtoDomain(menuItemDTO), nil
}
func (r *menuRepository) Create(ctx context.Context, menuItem *menu.Item) error {
	menuItemDTO := DomainToDTO(menuItem)
	_, err := r.db.Exec(ctx, `INSERT INTO menu (id, name, description, price, itemSize, allergens, categories, customization, ingredients) values ($1, $2, $3, $4,$5,$6,$7)`, menuItemDTO.Id, menuItemDTO.Name, menuItemDTO.Description, menuItemDTO.Price, menuItemDTO.ItemSize, menuItemDTO.Allergens, menuItemDTO.Categories)
	if err != nil {
		return err
	}

	return nil
}
func (r *menuRepository) DeleteById(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM menu WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
func (r *menuRepository) UpdateById(ctx context.Context, id string, newMenuItem *menu.Item) error {
	menuItemDTO := DomainToDTO(newMenuItem)
	_, err := r.db.Exec(ctx, `UPDATE menu SET name = $2, description = $3, price = $4, itemSize = $5, allergens = $6, categories = $7, customization = $8, ingredients = $9 WHERE id = $1`, id, menuItemDTO.Name, menuItemDTO.Description, menuItemDTO.Price, menuItemDTO.ItemSize, menuItemDTO.Allergens, menuItemDTO.Categories, menuItemDTO.Customization, menuItemDTO.Ingredients)
	if err != nil {
		return err
	}
	return nil
}
