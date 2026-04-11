package queries

type GetByIdResponse struct {
	Id            int
	Name          string
	Description   string
	Price         float64
	ItemSize      string
	Allergens     []string
	Categories    []string
	Customization string
	Ingredients   []ItemIngredient
}

type ItemIngredient struct {
	IngredientID string
	Quantity     float64
}
