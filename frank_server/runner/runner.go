package runner

import (
	"frank_server/models"
	"frank_server/scraper"
	"frank_server/scraper/allrecipes"
	"frank_server/utils"
)

// Runner is the main runner
type Runner struct {
	recipeName      string
	numberOfRecipes int
	searchScraper   scraper.SearchScraper
	recipeScraper   scraper.RecipeScraper
}

// NewRunner returns a new Runner
func NewRunner(recipeName string, numberOfRecipes int,
	searchScraper scraper.SearchScraper, recipeScraper scraper.RecipeScraper) Runner {
	return Runner{
		recipeName:      recipeName,
		numberOfRecipes: numberOfRecipes,
		searchScraper:   searchScraper,
		recipeScraper:   recipeScraper,
	}
}

// Run searches for receipes
func (r *Runner) Run() []*models.Recipe {
	searchRunner := SearchRunner{recipeName: utils.UrlFormat(r.recipeName), scraper: allrecipes.SearchScraper{}}

	// get links
	recipeURLs := searchRunner.Run()

	// get each recipe
	return r.fetchRecipes(recipeURLs)
}

func (r *Runner) fetchRecipes(recipeURLs []string) []*models.Recipe {
	recipes := []*models.Recipe{}

	for i := 0; i < r.numberOfRecipes; i++ {
		recipeRunner := NewRecipeRunner(recipeURLs[i], allrecipes.RecipeScraper{})
		result := recipeRunner.Run()
		recipes = append(recipes, result)
	}
	return recipes
}
