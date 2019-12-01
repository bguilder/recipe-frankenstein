package scraper

import "github.com/gocolly/colly"

// RecipeScraper comment
type RecipeScraper interface {
	IsIngredient(*colly.HTMLElement) bool
	IsTitle(*colly.HTMLElement) bool
	IsDirection(*colly.HTMLElement) bool
}

// SearchScraper comment
type SearchScraper interface {
	ShouldVisitLink(*colly.HTMLElement) bool
}
