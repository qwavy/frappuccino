package menu

type Item struct {
	id            int
	name          string
	description   string
	price         float64
	itemSize      string
	allergens     []string
	categories    []string
	customization string
	ingredients   []ItemIngredient
}

func NewMenuItem(
	id int,
	name string,
	description string,
	price float64,
	itemSize string,
	allergens []string,
	categories []string,
	customization string,
	ingredients []ItemIngredient,
) (*Item, error) {
	return &Item{
		id:            id,
		name:          name,
		description:   description,
		price:         price,
		itemSize:      itemSize,
		allergens:     allergens,
		categories:    categories,
		customization: customization,
		ingredients:   ingredients,
	}, nil
}

func RestoreMenuItem(id int,
	name string,
	description string,
	price float64,
	itemSize string,
	allergens []string,
	categories []string,
	customization string,
	ingredients []ItemIngredient) *Item {
	return &Item{
		id:            id,
		name:          name,
		description:   description,
		price:         price,
		itemSize:      itemSize,
		allergens:     allergens,
		categories:    categories,
		customization: customization,
		ingredients:   ingredients,
	}
}

func (i *Item) Id() int {
	return i.id
}
func (i *Item) Name() string {
	return i.name
}
func (i *Item) Description() string {
	return i.description
}
func (i *Item) Price() float64 {
	return i.price
}
func (i *Item) ItemSize() string {
	return i.itemSize
}
func (i *Item) Allergens() []string {
	return i.allergens
}
func (i *Item) Categories() []string {
	return i.categories
}
func (i *Item) Customization() string {
	return i.customization
}
func (i *Item) Ingredients() []ItemIngredient {
	return i.ingredients
}
