package menu_repo

type MenuItemDTO struct {
	Id            int                     `json:"product_id"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	Price         float64                 `json:"price"`
	ItemSize      string                  `json:"item_size"`
	Allergens     []string                `json:"allergens"`
	Categories    []string                `json:"categories"`
	Customization string                  `json:"customization"`
	Ingredients   []MenuItemIngredientDTO `json:"ingredients"`
}

type MenuItemIngredientDTO struct {
	IngredientID string  `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
}
