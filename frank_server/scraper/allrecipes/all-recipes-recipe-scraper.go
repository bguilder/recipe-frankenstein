package allrecipes

import (
	"frank_server/scraper"
	"frank_server/utils"

	"github.com/gocolly/colly"
)

const mainSelectorRecipe = "li, h1, span"

// RecipeScraper comment
type RecipeScraper struct {
}

// TryGetIngredient comment
func (s RecipeScraper) TryGetIngredient(e *colly.HTMLElement) string {
	if e.DOM.HasClass("ingredients-item-name") {
		return utils.NormalizeString(e.Text)
	} else if e.DOM.HasClass("checkList__line") {
		return utils.NormalizeString(e.Text)
	}
	return ""
}

// TryGetTitle comment
func (s RecipeScraper) TryGetTitle(e *colly.HTMLElement) string {
	if e.Name == "h1" {
		return utils.NormalizeString(e.Text)
	}
	return ""
}

// TryGetDirection comment
func (s RecipeScraper) TryGetDirection(e *colly.HTMLElement) string {
	if e.DOM.HasClass("recipe-directions__list--item") {
		return utils.NormalizeString(e.Text)
	} else if e.DOM.HasClass("instructions-section-item") {
		return utils.NormalizeString(e.Text)
	}
	return ""
}

// GetConfig comment
func (s RecipeScraper) GetConfig() scraper.RecipeScraperConfig {
	return scraper.RecipeScraperConfig{
		MainSelector: mainSelectorRecipe,
	}
}
