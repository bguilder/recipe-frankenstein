package models

import "frank_server/postprocessor"

// Recipe comment
type Recipe struct {
	Title       string
	Ingredients []string
	Directions  []string
}

type RecipesView struct {
	Recipes     []*Recipe
	Ingredients postprocessor.PairList
}

// AppendIngredient comment
func (r *Recipe) AppendIngredient(ingredient string) {
	r.Ingredients = append(r.Ingredients, ingredient)
}

// AppendDirection comment
func (r *Recipe) AppendDirection(direction string) {
	r.Directions = append(r.Directions, direction)
}

// SetTitle comment
func (r *Recipe) SetTitle(title string) {
	if r.Title != "" {
		return
	}
	r.Title = title
}
