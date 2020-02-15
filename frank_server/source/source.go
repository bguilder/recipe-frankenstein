package source

import "github.com/gocolly/colly"

// ILink defines the functionality of ILink source
type ILink interface {

	// returns either a link or empty string if no link is found
	TryGetRecipeLink(*colly.HTMLElement) string

	// returns the source specific link config
	GetConfig() LinkSourceConfig
}

// IRecipe defines the functionality of IRecipe source
type IRecipe interface {

	// returns an ingredient as a string or an empty string if none is found
	TryGetIngredient(*colly.HTMLElement) string

	// returns the recipe title as a string or an empty string if none is found
	TryGetTitle(*colly.HTMLElement) string

	// return a direction as a string or an empty string if none is found
	TryGetDirection(*colly.HTMLElement) string

	// returns the source specific recipe config
	GetConfig() RecipeSourceConfig
}
