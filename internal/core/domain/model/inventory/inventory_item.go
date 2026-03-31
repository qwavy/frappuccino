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
