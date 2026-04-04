package order

type Item struct {
	productID string
	quantity  int
}

func NewOrderItem(productID string, quantity int) *Item {
	return &Item{productID: productID, quantity: quantity}
}

func RestoreOrderItem(productID string, quantity int) *Item {
	return &Item{productID: productID, quantity: quantity}
}

func (i *Item) ProductID() string {
	return i.productID
}
func (i *Item) Quantity() int {
	return i.quantity
}
