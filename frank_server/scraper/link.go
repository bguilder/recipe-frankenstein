package scraper

import (
	"fmt"
	"frank_server/source"
	"frank_server/utils"

	"github.com/gocolly/colly"
)

// ILinkScraper defines the functionality of a link scraper
type ILinkScraper interface {
	// returns a slice of recipe urls
	Run(recipeName string, recipeCount int) []string
}

// linkScraper comment
type linkScraper struct {
	links  []string
	source source.ILink
}

// NewLinkScraper constructor returns an initialized ILinkScraper
func NewLinkScraper(source source.ILink) ILinkScraper {
	return &linkScraper{source: source, links: make([]string, 0, 5)}
}

func (s *linkScraper) appendRecipeLink(link string) {
	if link != "" {
		s.links = append(s.links, link)
	}
}

// Run comment
func (s *linkScraper) run(recipeName string, recipeCount int) {
	c := colly.NewCollector()
	config := s.source.GetConfig()
	fmt.Printf("name: %+v", recipeName)

	c.OnHTML(config.MainSelector, func(e *colly.HTMLElement) {
		if len(s.links) >= recipeCount {
			s.logOnScraped()
			return
		}
		if link := s.source.TryGetRecipeLink(e); link != "" {
			s.appendRecipeLink(link)
		}
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	s.logOnScraped()
	// })

	// Start scraping
	c.Visit(config.Domain + config.LinkPath + utils.UrlFormat(recipeName) + "&sort=re")
}

// Run satisfies the source.ILinkScraper interface
func (s *linkScraper) Run(recipeName string, recipeCount int) []string {
	s.run(recipeName, recipeCount)
	return s.links
}

func (s *linkScraper) logOnScraped() {
	fmt.Printf("\n\n\n============Links============\n\n\n")
	fmt.Printf("Links: %s\n\n", s.links)
}
