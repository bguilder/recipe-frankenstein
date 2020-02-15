package main

import (
	"fmt"
	"frank_server/runner"
	"frank_server/scraper"
	"frank_server/source/allrecipes"
)

func main() {
	runner := runner.NewRunner(
		"chicken parmesan",
		2,
		scraper.NewLinkScraper(&allrecipes.LinkSource{}),
		scraper.NewRecipeScraper(&allrecipes.RecipeSource{}))
	recipes := runner.Run()
	for _, recipe := range recipes {
		fmt.Printf("recipe: %+v\n\n", recipe)
	}
}
