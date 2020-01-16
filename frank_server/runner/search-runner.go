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
	recipeLinks []string
	recipeName  string
	scraper     scraper.SearchScraper
}

func (s *SearchRunner) AppendRecipeLink(link string) {
	if link != "" {
		s.recipeLinks = append(s.recipeLinks, link)
	}
}

// Run comment
func (s *SearchRunner) run() {
	c := colly.NewCollector()
	config := s.scraper.GetConfig()

	c.OnHTML(config.MainSelector, func(e *colly.HTMLElement) {
		if link := s.scraper.TryGetRecipeLink(e); link != "" {
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
	c.Visit(config.Domain + config.SearchPath + s.recipeName + "&sort=re")
}

func (s *SearchRunner) Run() []string {
	s.run()
	return s.recipeLinks
}
