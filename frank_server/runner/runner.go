package runner

import (
	"frank_server/models"
	"frank_server/scraper/allrecipes"
	"frank_server/utils"
	"time"
)

type Runner struct {
}

func Run(recipeName string, numberOfRecipes int) []string {
	searchRunner := SearchRunner{RecipeName: utils.UrlFormat(recipeName)}
	SearchScraper := allrecipes.SearchScraper{}

	searchRunner.Run(SearchScraper)

	recipes := []*models.Recipe{}

	for i := 0; i < numberOfRecipes; i++ {
		recipe := models.Recipe{}
		recipeRunner := RecipeRunner{Recipe: recipe, RecipeLink: searchRunner.RecipeLinks[i]}
		RecipeScraper := allrecipes.RecipeScraper{}
		recipeRunner.Run(RecipeScraper)
		recipes = append(recipes, &recipeRunner.Recipe)
		time.Sleep(2 * time.Second)
	}

	allIngredients := []string{}
	for i := 0; i < len(recipes); i++ {
		for x := 0; x < len(recipes[i].Ingredients); x++ {
			allIngredients = append(allIngredients, recipes[i].Ingredients[x])
		}
	}
	return allIngredients
}
