package allrecipes

import (
	"frank_server/scraper"
	"frank_server/utils"

	"github.com/gocolly/colly"
)

const mainSelectorSearch = "h3"
const domain = "https://www.allrecipes.com"
const searchPath = "/search/results/?wt="

// SearchScraper comment
type SearchScraper struct {
}

// TryGetRecipeLink comment
func (s SearchScraper) TryGetRecipeLink(e *colly.HTMLElement) string {
	// TODO: Add additional check for the correct title
	if e.DOM.HasClass("fixed-recipe-card__h3") {
		recipeLink := e.ChildAttr("*", "href")
		return utils.NormalizeString(recipeLink)
	}
	return ""
}

// GetConfig comment
func (s SearchScraper) GetConfig() scraper.SearchScraperConfig {
	return scraper.SearchScraperConfig{
		MainSelector: mainSelectorSearch,
		Domain:       domain,
		SearchPath:   searchPath,
	}
}
