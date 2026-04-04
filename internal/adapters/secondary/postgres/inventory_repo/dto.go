package inventory_repo

import "time"

type InventoryItemDTO struct {
	IngredientID int       `json:"ingredient_id"`
	Name         string    `json:"name"`
	Quantity     float64   `json:"quantity"`
	Unit         string    `json:"unit"`
	Price        float64   `json:"price"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
}
