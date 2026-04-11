package queries

import "time"

type Item struct {
	IngredientID int
	Name         string
	Quantity     float64
	Unit         string
	Price        float64
	Created      time.Time
	Updated      time.Time
}

type GetAllQueryResponse struct {
	InventoryItems []Item
}
