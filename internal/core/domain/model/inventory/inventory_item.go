package inventory

import (
	"errors"
	"time"
)

type Item struct {
	ingredientID int
	name         string
	quantity     float64
	unit         Unit
	price        float64
	created      time.Time
	updated      time.Time
}

func NewInventoryItem(ingredientID int, name string, quantity float64, unit string, price float64) (*Item, error) {
	myUnit, err := NewUnit(unit)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return &Item{
		ingredientID: ingredientID,
		name:         name,
		quantity:     quantity,
		unit:         myUnit,
		price:        price,
		created:      now,
		updated:      now,
	}, nil
}

func RestoreInventoryItem(ingredientID int, name string, quantity float64, unit Unit, price float64, created time.Time, updated time.Time) *Item {
	return &Item{
		ingredientID: ingredientID,
		name:         name,
		quantity:     quantity,
		unit:         unit,
		price:        price,
		created:      created,
		updated:      updated,
	}
}

func (i *Item) IngredientID() int {
	return i.ingredientID
}
func (i *Item) Name() string {
	return i.name
}
func (i *Item) Quantity() float64 {
	return i.quantity
}
func (i *Item) Unit() Unit {
	return i.unit
}
func (i *Item) Price() float64 {
	return i.price
}
func (i *Item) Created() time.Time {
	return i.created
}
func (i *Item) Updated() time.Time {
	return i.updated
}

func (i *Item) Update(name string, quantity float64, unit Unit, price float64) {
	now := time.Now().UTC()

	i.name = name
	i.quantity = quantity
	i.unit = unit
	i.price = price
	i.updated = now
}

func (i *Item) DecreaseQuantity(amount float64) error {
	if amount < 0 {
		return errors.New("given param can`t be negative")
	}

	if (i.quantity - amount) == 0 {
		return errors.New("quantity cant be negative")
	}

	i.quantity -= amount

	return nil
}
