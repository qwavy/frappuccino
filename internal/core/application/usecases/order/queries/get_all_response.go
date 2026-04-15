package queries

import "time"

type GetAllResponse struct {
	orders []GetAllItem
}

type GetAllItem struct {
	Id           string
	CustomerName string
	Items        []OrderItem
	Status       string
	CreatedAt    time.Time
}

type OrderItem struct {
	ProductID string
	Quantity  int
}
