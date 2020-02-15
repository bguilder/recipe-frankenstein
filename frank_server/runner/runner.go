package runner

import (
	"fmt"
	"frank_server/scraper"
)

// Runner is the main runner
type Runner struct {
	recipeName      string
	numberOfRecipes int
	linkScraper     scraper.ILinkScraper
	recipeScraper   scraper.IRecipeScraper
}

// NewRunner returns a new Runner
func NewRunner(recipeName string, numberOfRecipes int,
	linkScraper scraper.ILinkScraper, recipeScraper scraper.IRecipeScraper) Runner {
	return Runner{
		recipeName:      recipeName,
		numberOfRecipes: numberOfRecipes,
		linkScraper:     linkScraper,
		recipeScraper:   recipeScraper,
	}
}

// Run searches for receipes
func (r *Runner) Run() []*scraper.Recipe {

	recipeURLs := r.linkScraper.Run(r.recipeName, r.numberOfRecipes)
	fmt.Printf("recipeURLs: %+v", recipeURLs)
	// get each recipe
	return r.fetchRecipes(recipeURLs)
}

func (r *Runner) fetchRecipes(recipeURLs []string) []*scraper.Recipe {
	recipes := []*scraper.Recipe{}

	for i := 0; i < len(recipeURLs); i++ {
		result := r.recipeScraper.Run(recipeURLs[i])
		recipes = append(recipes, result)
	}
	return recipes
}
