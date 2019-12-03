package runner

import (
	"fmt"
	"frank_server/scraper"
	"frank_server/utils"
	"log"

	"github.com/gocolly/colly"
)

// SearchRunner comment
type SearchRunner struct {
	RecipeLinks []string
	RecipeName  string
}

func (s *SearchRunner) AppendRecipeLink(link string) {
	if link != "" {
		s.RecipeLinks = append(s.RecipeLinks, link)
	}
}

// Run comment
func (s *SearchRunner) Run(scraper scraper.SearchScraper) {
	c := colly.NewCollector()
	config := scraper.GetConfig()

	c.OnHTML(config.MainSelector, func(e *colly.HTMLElement) {
		if link := scraper.TryGetRecipeLink(e); link != "" {
			s.AppendRecipeLink(link)
			fmt.Printf("Recipe Name: %s\n Link: %s\n", utils.NormalizeString(e.Text), e.ChildAttr("*", "href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Print("\n\n======Recipes List=======\n")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Finished!\n\n")
	})

	// Start scraping
	c.Visit(config.Domain + config.SearchPath + s.RecipeName + "&sort=re")
}
