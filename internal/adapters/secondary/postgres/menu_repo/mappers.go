package menu_repo

import "frappucchino/internal/core/domain/model/menu"

func DomainToDTO(item *menu.Item) MenuItemDTO {
	var menuItemDTO MenuItemDTO

	menuItemDTO.Id = item.Id()
	menuItemDTO.Name = item.Name()
	menuItemDTO.Description = item.Description()
	menuItemDTO.Price = item.Price()
	menuItemDTO.ItemSize = item.ItemSize()
	menuItemDTO.Allergens = item.Allergens()
	menuItemDTO.Categories = item.Categories()
	menuItemDTO.Customization = item.Customization()

	var ingredientsDTO []MenuItemIngredientDTO

	for _, ingredient := range item.Ingredients() {
		var ingredientDTO MenuItemIngredientDTO
		ingredientDTO.IngredientID = ingredient.IngredientID()
		ingredientDTO.Quantity = ingredient.Quantity()

		ingredientsDTO = append(ingredientsDTO, ingredientDTO)
	}

	menuItemDTO.Ingredients = ingredientsDTO

	return menuItemDTO
}

func DTOtoDomain(menuItemDTO MenuItemDTO) *menu.Item {
	aggregate := menu.RestoreMenuItem(menuItemDTO.Id, menuItemDTO.Name, menuItemDTO.Description, menuItemDTO.Price, menuItemDTO.ItemSize, menuItemDTO.Allergens, menuItemDTO.Categories, menuItemDTO.Customization, menuItemDTO.Ingredients)
	return aggregate
}
