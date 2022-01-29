package runner

import (
	"frank_server/db"
	"frank_server/models"
	"frank_server/postprocessor"
	"log"
	"strings"
)

const defaultRecipeCount = 7

// SearchRunner is the main runner
type SearchRunner struct {
	scraper Scraper
	cache   db.Store
}

// NewSearchRunner returns a new SearchRunner
func NewSearchRunner(cache db.Store, scraper Scraper) SearchRunner {
	return SearchRunner{
		scraper: scraper,
		cache:   cache,
	}
}

// Run searches for receipes
func (r *SearchRunner) Run(recipeName string, count int) models.RecipesView {
	formattedRecipeName := strings.ToLower(recipeName)

	// Set default max search recipes
	if count > defaultRecipeCount || count <= 0 {
		count = defaultRecipeCount
	}

	recipes, err := r.cache.GetRecipes(formattedRecipeName)
	if err != nil {
		log.Panic(err)
	}

	if recipes != nil {
		log.Println("loaded from cache")
		return buildRecipesView(recipes)
	}

	log.Println("cache miss...")

	result := r.scraper.GetRecipes(formattedRecipeName, count)

	// update cache
	err = r.cache.PutRecipes(formattedRecipeName, result)
	if err != nil {
		log.Panic(err)
	}
	return buildRecipesView(recipes)
}

func buildRecipesView(recipes []*models.Recipe) models.RecipesView {
	ing := formatIngredients(recipes)
	log.Printf("formatted ingredients: %+v\n", ing)

	return models.RecipesView{Recipes: recipes, Ingredients: ing}
}

func formatIngredients(recipes []*models.Recipe) postprocessor.IngredientFrequencyList {
	pp := postprocessor.NewPostProcessor()

	ingredients := []string{}
	for _, recipe := range recipes {
		ingredients = append(ingredients, recipe.Ingredients...)
	}

	return pp.Run(ingredients)
}
