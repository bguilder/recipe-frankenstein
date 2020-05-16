package api

import (
	"frank_server/cache"
	"frank_server/postprocessor"
	"frank_server/runner"
	"frank_server/scraper"
	"frank_server/source/allrecipes"
	"log"
)

type RecipeFetcherService struct {
	store       cache.Store
	recipeName  string
	recipeCount int
}

func NewRecipeFetcherService(store cache.Store, recipeName string, recipeCount int) RecipeFetcherService {
	return RecipeFetcherService{store: store, recipeName: recipeName, recipeCount: recipeCount}
}

func (r RecipeFetcherService) Run() RecipesView {
	// try to fetch from cache
	recipes, err := r.store.GetRecipes(r.recipeName)
	if err != nil {
		log.Panic(err)
	}
	if recipes != nil {
		log.Println("loaded from cache")
		return NewRecipesView(recipes, formatIngredients(recipes))
	}

	log.Println("cache miss...")

	runner := runner.NewRunner(
		r.recipeName,
		r.recipeCount,
		scraper.NewLinkScraper(&allrecipes.LinkSource{}),
		scraper.NewRecipeScraper(&allrecipes.RecipeSource{}))

	recipes = runner.Run()

	// update cache
	err = r.store.PutRecipes(r.recipeName, recipes)
	if err != nil {
		log.Panic(err)
	}

	return NewRecipesView(recipes, formatIngredients(recipes))
}

func formatIngredients(recipes []*scraper.Recipe) postprocessor.IngredientFrequencyList {
	pp := postprocessor.NewPostProcessor()

	ingredients := []string{}
	for _, recipe := range recipes {
		for _, ing := range recipe.Ingredients {
			ingredients = append(ingredients, ing)
		}
	}

	return pp.Run(ingredients)
}
