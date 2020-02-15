package allrecipes

import (
	"frank_server/source"
	"frank_server/utils"

	"github.com/gocolly/colly"
)

const mainSelectorSearch = "h3"
const domain = "https://www.allrecipes.com"
const linkPath = "/search/results/?wt="

// LinkSource satisfies the source.ILink interface
type LinkSource struct {
}

// TryGetRecipeLink satisfies the source.ILink interface
func (s LinkSource) TryGetRecipeLink(e *colly.HTMLElement) string {
	// TODO: Add additional check for the correct title
	if e.DOM.HasClass("fixed-recipe-card__h3") {
		//fmt.Printf("Recipe Name: %s\n Link: %s\n", utils.NormalizeString(e.Text), e.ChildAttr("*", "href"))

		recipeLink := e.ChildAttr("*", "href")
		return utils.NormalizeString(recipeLink)
	}
	return ""
}

// GetConfig satisfies the source.ILink interface
func (s LinkSource) GetConfig() source.LinkSourceConfig {
	return source.LinkSourceConfig{
		MainSelector: mainSelectorSearch,
		Domain:       domain,
		LinkPath:     linkPath,
	}
}
