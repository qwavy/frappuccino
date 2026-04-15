package queries

import "time"

type GetByIdResponse struct {
	Id           string
	CustomerName string
	Items        []OrderItem
	Status       string
	CreatedAt    time.Time
}
