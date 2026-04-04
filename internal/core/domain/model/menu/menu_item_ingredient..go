package menu

import "errors"

type ItemIngredient struct {
	ingredientID string
	quantity     float64
}

func NewItemIngredient(IngredientID string, quantity float64) (ItemIngredient, error) {
	if IngredientID == "" {
		return ItemIngredient{}, errors.New("ingredient ID is required")
	}
	if quantity < 1 {
		return ItemIngredient{}, errors.New("quantity cant be negative")
	}

	return ItemIngredient{ingredientID: IngredientID, quantity: quantity}, nil
}

func (i *ItemIngredient) IngredientID() string {
	return i.ingredientID
}

func (i *ItemIngredient) Quantity() float64 {
	return i.quantity
}
