package allrecipes

import (
	"frank_server/source"
	"frank_server/utils"

	"github.com/gocolly/colly"
)

const mainSelectorRecipe = "li, h1, span"

// RecipeSource satisfies the source.IRecipe interface
type RecipeSource struct {
}

// TryGetIngredient satisfies the source.IRecipe interface
func (s RecipeSource) TryGetIngredient(e *colly.HTMLElement) string {
	if e.DOM.HasClass("ingredients-item-name") {
		return utils.NormalizeString(e.Text)
	} else if e.DOM.HasClass("checkList__line") {
		return utils.NormalizeString(e.Text)
	}
	return ""
}

// TryGetTitle satisfies the source.IRecipe interface
func (s RecipeSource) TryGetTitle(e *colly.HTMLElement) string {
	if e.Name == "h1" {
		return utils.NormalizeString(e.Text)
	}
	return ""
}

// TryGetDirection satisfies the source.IRecipe interface
func (s RecipeSource) TryGetDirection(e *colly.HTMLElement) string {
	if e.DOM.HasClass("recipe-directions__list--item") {
		return utils.NormalizeString(e.Text)
	} else if e.DOM.HasClass("instructions-section-item") {
		return utils.NormalizeString(e.Text)
	}
	return ""
}

// GetConfig satisfies the source.IRecipe interface
func (s RecipeSource) GetConfig() source.RecipeSourceConfig {
	return source.RecipeSourceConfig{
		MainSelector: mainSelectorRecipe,
	}
}
