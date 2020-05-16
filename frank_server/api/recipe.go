package api

import (
	"encoding/json"
	"frank_server/postprocessor"
	"frank_server/scraper"
)

// RecipesView is the view object that is sent to the client
type RecipesView struct {
	Recipes     []*scraper.Recipe
	Ingredients postprocessor.IngredientFrequencyList
}

// NewRecipesView returns a new instance of a RecipesView
func NewRecipesView(recipes []*scraper.Recipe, ingredients postprocessor.IngredientFrequencyList) RecipesView {
	return RecipesView{Recipes: recipes, Ingredients: ingredients}
}

// ToJSONBytes returns the object as json bytes
func (r *RecipesView) ToJSONBytes() ([]byte, error) {
	return json.Marshal(r)
}

// ToJSONString returns the object as a json string
func (r *RecipesView) ToJSONString() (string, error) {
	payload, err := json.Marshal(r)
	return string(payload), err
}
