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
	Recipe     models.Recipe
	RecipeLink string
}

// Run comment
func (s *RecipeRunner) Run(scraper scraper.RecipeScraper) {
	c := colly.NewCollector()

	c.OnHTML(scraper.GetConfig().MainSelector, func(e *colly.HTMLElement) {

		// Try to add title
		if title := scraper.TryGetTitle(e); title != "" {
			s.Recipe.SetTitle(title)
			return
		}

		// Try to add ingredients
		if ingredient := scraper.TryGetIngredient(e); ingredient != "" {
			s.Recipe.AppendIngredient(ingredient)
			return
		}

		// Try to add directions
		if direction := scraper.TryGetDirection(e); direction != "" {
			s.Recipe.AppendDirection(direction)
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
		fmt.Printf("Title: %s\n\n", s.Recipe.Title)
		fmt.Printf("Directions\n\n")
		for i := 0; i < len(s.Recipe.Directions); i++ {
			fmt.Printf("Step %v: %s\n", i+1, s.Recipe.Directions[i])
		}
		fmt.Printf("Ingredients\n\n")
		for i := 0; i < len(s.Recipe.Ingredients); i++ {
			fmt.Printf("%s\n", s.Recipe.Ingredients[i])
		}
	})

	// Start scraping
	c.Visit(s.RecipeLink)
}
