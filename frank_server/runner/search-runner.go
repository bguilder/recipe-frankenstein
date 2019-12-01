package runner

import (
	"fmt"
	"frank_server/scraper"
	"log"

	"github.com/gocolly/colly"
)

// SearchRunner comment
type SearchRunner struct {
	RecipeLinks []string
	RecipeName  string
}

// Run comment
func (s *SearchRunner) Run(scraper scraper.SearchScraper) {
	c := colly.NewCollector()

	c.OnHTML(mainSelectorSearch, func(e *colly.HTMLElement) {
		if scraper.ShouldVisitLink(e) {
			recipeLink := e.ChildAttr("*", "href")
			s.RecipeLinks = append(s.RecipeLinks, recipeLink)
			fmt.Printf("Recipe Name: %s\n Link: %s\n", normalizeString(e.Text), e.ChildAttr("*", "href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("\n\n======Recipes List=======\n")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Finished!\n\n")
	})

	// Start scraping
	c.Visit(domain + searchPath + s.RecipeName + "&sort=re")
}
