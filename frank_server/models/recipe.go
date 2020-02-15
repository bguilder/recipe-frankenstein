package models

import (
	"frank_server/postprocessor"
	"frank_server/scraper"
)

type RecipesView struct {
	Recipes     []*scraper.Recipe
	Ingredients postprocessor.PairList
}
