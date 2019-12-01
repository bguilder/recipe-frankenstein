package scraper

import "github.com/gocolly/colly"

// AllRecipesRecipeScraper comment
type AllRecipesRecipeScraper struct {
}

// IsIngredient comment
func (s AllRecipesRecipeScraper) IsIngredient(e *colly.HTMLElement) bool {
	if e.DOM.HasClass("ingredients-item-name") {
		return true
	} else if e.DOM.HasClass("checkList__line") {
		return true
	}
	return false
}

// IsTitle comment
func (s AllRecipesRecipeScraper) IsTitle(e *colly.HTMLElement) bool {
	if e.Name == "h1" {
		return true
	}
	return false
}

// IsDirection comment
func (s AllRecipesRecipeScraper) IsDirection(e *colly.HTMLElement) bool {
	if e.DOM.HasClass("recipe-directions__list--item") {
		return true
	} else if e.DOM.HasClass("instructions-section-item") {
		return true
	}
	return false
}
