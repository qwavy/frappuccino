package reports

type SearchMenuItemReport struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       string  `json:"price"`
	Relevance   float64 `json:"relevance"`
}

type SearchOrderReport struct {
	ID           int      `json:"id"`
	CustomerName string   `json:"customer_name"`
	Items        []string `json:"description"`
	Price        string   `json:"total"`
	Relevance    float64  `json:"relevance"`
}

type SearchReport struct {
	MenuItems    []*SearchMenuItemReport `json:"menu_items,omitempty"`
	Orders       []*SearchOrderReport    `json:"orders,omitempty"`
	TotalMatches int                     `json:"total_matches"`
}

type InventoryItemReport struct {
	Name        string  `json:"name"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	TotalAmount int     `json:"-"`
}

type LeftOvers struct {
	CurrentPage int                   `json:"current_page"`
	HasNextPage bool                  `json:"hasNextPage"`
	PageSize    int                   `json:"pageSize"`
	TotalPages  int                   `json:"totalPages"`
	Items       []InventoryItemReport `json:"data"`
}
