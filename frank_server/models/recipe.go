package models

// Recipe model contains all scraped recipe data
type Recipe struct {
	Title       string
	Ingredients []string
	Directions  []string
	URL         string
}

func (r *Recipe) AppendIngredient(ingredient string) {
	r.Ingredients = append(r.Ingredients, ingredient)
}

func (r *Recipe) AppendDirection(direction string) {
	r.Directions = append(r.Directions, direction)
}
