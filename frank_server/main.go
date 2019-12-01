package main

import (
	"fmt"
	"frank_server/models"
	"frank_server/runner"
	"frank_server/scraper"
	"time"
)

const recipeName = "christmas cookies"
const numberOfRecipes = 3

func main() {
	fmt.Printf("Searching AllRecipes for: %s\n\n", recipeName)
	searchRunner := runner.SearchRunner{RecipeName: runner.UrlFormat(recipeName)}
	allRecipesSearchScraper := scraper.AllRecipesSearchScraper{}

	searchRunner.Run(allRecipesSearchScraper)

	recipes := []*models.Recipe{}

	for i := 0; i < numberOfRecipes; i++ {
		recipe := models.Recipe{}
		recipeRunner := runner.RecipeRunner{Recipe: recipe, RecipeLink: searchRunner.RecipeLinks[i]}
		allRecipesRecipeScraper := scraper.AllRecipesRecipeScraper{}
		recipeRunner.Run(allRecipesRecipeScraper)
		recipes = append(recipes, &recipeRunner.Recipe)
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\n\n\n=================================Total ingredients list!=================================\n\n\n")
	for i := 0; i < len(recipes); i++ {
		for x := 0; x < len(recipes[i].Ingredients)-1; x++ {
			fmt.Printf("- %s\n", recipes[i].Ingredients[x])
		}
	}
}
