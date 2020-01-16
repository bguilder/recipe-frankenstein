package runner

import (
	"fmt"
	"frank_server/models"
	"frank_server/scraper"
	"log"

	"github.com/gocolly/colly"
)

// RecipeRunner comment
type RecipeRunner struct {
	recipe     models.Recipe
	recipeLink string
	scraper    scraper.RecipeScraper
}

func NewRecipeRunner(recipeURL string, scraper scraper.RecipeScraper) RecipeRunner {
	return RecipeRunner{recipe: models.Recipe{}, recipeLink: recipeURL, scraper: scraper}
}

// Run comment
func (s *RecipeRunner) run() {
	c := colly.NewCollector()

	c.OnHTML(s.scraper.GetConfig().MainSelector, func(e *colly.HTMLElement) {

		// Try to add title
		if title := s.scraper.TryGetTitle(e); title != "" {
			s.recipe.SetTitle(title)
			return
		}

		// Try to add ingredients
		if ingredient := s.scraper.TryGetIngredient(e); ingredient != "" {
			s.recipe.AppendIngredient(ingredient)
			return
		}

		// Try to add directions
		if direction := s.scraper.TryGetDirection(e); direction != "" {
			s.recipe.AppendDirection(direction)
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("\n\n\n============Recipe============\n\n\n")
		fmt.Printf("Title: %s\n\n", s.recipe.Title)
		fmt.Printf("Directions\n\n")
		for i := 0; i < len(s.recipe.Directions); i++ {
			fmt.Printf("Step %v: %s\n", i+1, s.recipe.Directions[i])
		}
		fmt.Printf("Ingredients\n\n")
		for i := 0; i < len(s.recipe.Ingredients); i++ {
			fmt.Printf("%s\n", s.recipe.Ingredients[i])
		}
	})

	// Start scraping
	c.Visit(s.recipeLink)
}

func (s *RecipeRunner) Run() *models.Recipe {
	s.run()
	return &s.recipe
}
