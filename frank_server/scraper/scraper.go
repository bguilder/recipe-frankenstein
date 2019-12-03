package scraper

import "github.com/gocolly/colly"

type RecipeScraperConfig struct {
	MainSelector string
}

// RecipeScraper comment
type RecipeScraper interface {
	TryGetIngredient(*colly.HTMLElement) string
	TryGetTitle(*colly.HTMLElement) string
	TryGetDirection(*colly.HTMLElement) string
	GetConfig() RecipeScraperConfig
}

type SearchScraperConfig struct {
	MainSelector string
	SearchPath   string
	Domain       string
}

// SearchScraper comment
type SearchScraper interface {
	TryGetRecipeLink(*colly.HTMLElement) string
	GetConfig() SearchScraperConfig
}
