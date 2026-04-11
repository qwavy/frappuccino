package order

import (
	"errors"
	"time"
)

var (
	ErrOrderClosed = errors.New("order already closed")
)

type Order struct {
	id           string
	customerName string
	items        []Item
	status       Status
	createdAt    time.Time
}

func NewOrder(id string, customerName string, items []Item, status Status) (*Order, error) {
	now := time.Now().UTC()

	return &Order{
		id:           id,
		customerName: customerName,
		items:        items,
		status:       StatusOpened,
		createdAt:    now,
	}, nil
}

func RestoreOrder(id string, customerName string, items []Item, status Status, createdAt time.Time) *Order {
	return &Order{
		id:           id,
		customerName: customerName,
		items:        items,
		status:       status,
		createdAt:    createdAt,
	}
}

func (o *Order) ID() string {
	return o.id
}

func (o *Order) CustomerName() string {
	return o.customerName
}
func (o *Order) Items() []Item {
	return o.items
}
func (o *Order) Status() Status {
	return o.status
}
func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) Equals(anotherOrder *Order) bool {
	return o.id == anotherOrder.id
}

func (o *Order) Close() error {
	if o.status == StatusClosed {
		return ErrOrderClosed
	}

	o.status = StatusClosed

	return nil
}

func (o *Order) SetCustomerName(newName string) error {
	if newName == "" {
		return errors.New("new customer name cant be empty")
	}

	o.customerName = newName

	return nil
}
