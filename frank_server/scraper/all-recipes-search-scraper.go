package scraper

import "github.com/gocolly/colly"

// AllRecipesSearchScraper comment
type AllRecipesSearchScraper struct {
}

// ShouldVisitLink comment
func (s AllRecipesSearchScraper) ShouldVisitLink(e *colly.HTMLElement) bool {
	// TODO: Add additional check for the correct title
	if e.DOM.HasClass("fixed-recipe-card__h3") {
		return true
	}
	return false
}
