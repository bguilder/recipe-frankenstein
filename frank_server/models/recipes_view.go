package models

import (
	"encoding/json"
	"frank_server/postprocessor"
	"log"
)

type RecipesView struct {
	Recipes     []*Recipe
	Ingredients postprocessor.IngredientFrequencyList
}

func (r *RecipesView) Marshal() []byte {
	payload, err := json.Marshal(r)
	if err != nil {
		log.Panic(err)
	}
	return payload
}
