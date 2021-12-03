package runner

import "frank_server/models"

// Scraper defines the functionality of a scraper
type Scraper interface {
	GetRecipes(recipeName string, recipeCount int) []*models.Recipe
}
