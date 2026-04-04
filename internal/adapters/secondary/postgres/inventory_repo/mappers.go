package inventory_repo

import "frappucchino/internal/core/domain/model/inventory"

func DomainToDTO(aggregate *inventory.Item) InventoryItemDTO {
	var inventoryItemDTO InventoryItemDTO

	inventoryItemDTO.IngredientID = aggregate.IngredientID()
	inventoryItemDTO.Name = aggregate.Name()
	inventoryItemDTO.Quantity = aggregate.Quantity()
	inventoryItemDTO.Unit = aggregate.Unit().String()
	inventoryItemDTO.Price = aggregate.Price()
	inventoryItemDTO.Created = aggregate.Created()
	inventoryItemDTO.Updated = aggregate.Updated()

	return inventoryItemDTO
}

func DTOtoDomain(dto InventoryItemDTO) *inventory.Item {
	unit := inventory.Unit(dto.Unit)
	aggregate := inventory.RestoreInventoryItem(dto.IngredientID, dto.Name, dto.Quantity, unit, dto.Price, dto.Created, dto.Updated)

	return aggregate
}
