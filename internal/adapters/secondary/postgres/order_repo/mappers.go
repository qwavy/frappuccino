package order_repo

import "frappucchino/internal/core/domain/model/order"

func DomainToDTO(order *order.Order) OrderDTO {
	var orderDTO OrderDTO

	orderDTO.Id = order.ID()
	orderDTO.CustomerName = order.CustomerName()
	orderDTO.Status = order.Status().String()

	var orderItemsDTO []OrderItemDTO

	for _, orderItem := range order.Items() {
		var orderItemDTO OrderItemDTO
		orderItemDTO.ProductID = orderItem.ProductID()
		orderItemDTO.Quantity = orderItem.Quantity()

		orderItemsDTO = append(orderItemsDTO, orderItemDTO)
	}

	orderDTO.Items = orderItemsDTO
	orderDTO.CreatedAt = order.CreatedAt()

	return orderDTO
}

func DTOtoDomain(orderDTO OrderDTO) *order.Order {
	orderStatus := order.Status(orderDTO.Status)
	var orderItems []*order.Item

	for _, orderItem := range orderItems {
		orderItems = append(orderItems, orderItem)
	}

	order := order.RestoreOrder(orderDTO.Id, orderDTO.CustomerName, orderItems, orderStatus, orderDTO.CreatedAt)
	return order
}
