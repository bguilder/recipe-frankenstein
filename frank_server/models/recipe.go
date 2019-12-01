package models

// Recipe comment
type Recipe struct {
	Title       string
	Ingredients []string
	Directions  []string
}

func (r *Recipe) AppendIngredient(ingredient string) {
	if ingredient == "" {
		return
	}
	r.Ingredients = append(r.Ingredients, ingredient)
}

func (r *Recipe) AppendDirection(direction string) {
	r.Directions = append(r.Directions, direction)
}

func (r *Recipe) AddTitle(title string) {
	r.Title = title
}
