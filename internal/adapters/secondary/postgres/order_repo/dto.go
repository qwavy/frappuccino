package order_repo

import "time"

type OrderDTO struct {
	Id           string         `json:"id"`
	CustomerName string         `json:"customer_name"`
	Items        []OrderItemDTO `json:"items"`
	Status       string         `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
}

type OrderItemDTO struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
