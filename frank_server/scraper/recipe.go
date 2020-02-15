package scraper

import (
	"fmt"
	"frank_server/source"
	"log"

	"github.com/gocolly/colly"
)

// Recipe model contains all scraped recipe data
type Recipe struct {
	Title       string
	Ingredients []string
	Directions  []string
	URL         string
}

func (r *Recipe) appendIngredient(ingredient string) {
	r.Ingredients = append(r.Ingredients, ingredient)
}

func (r *Recipe) appendDirection(direction string) {
	r.Directions = append(r.Directions, direction)
}

func (r *Recipe) setTitle(title string) {
	if r.Title == "" {
		r.Title = title
	}
}

// IRecipeScraper defines the functionality of a recipe scraper
type IRecipeScraper interface {
	// returns a fully hydrated recipe
	Run(url string) *Recipe
}

type recipeScraper struct {
	recipe Recipe
	source source.IRecipe
}

// NewRecipeScraper constructor returns an initialized IRecipeScraper
func NewRecipeScraper(source source.IRecipe) IRecipeScraper {
	return &recipeScraper{recipe: Recipe{}, source: source}
}

func (s *recipeScraper) run(url string) {
	c := colly.NewCollector()

	s.recipe.URL = url

	c.OnHTML(s.source.GetConfig().MainSelector, func(e *colly.HTMLElement) {

		if title := s.source.TryGetTitle(e); title != "" {
			s.recipe.setTitle(title)
			return
		}

		if ingredient := s.source.TryGetIngredient(e); ingredient != "" {
			s.recipe.appendIngredient(ingredient)
			return
		}

		if direction := s.source.TryGetDirection(e); direction != "" {
			s.recipe.appendDirection(direction)
			return
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error scraping:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		s.logOnScraped()
	})

	// Start scraping
	c.Visit(s.recipe.URL)
}

// Run satisfies the source.IRecipeScraper interface
func (s *recipeScraper) Run(url string) *Recipe {
	s.run(url)
	result := s.recipe
	s.recipe = Recipe{}
	return &result
}

func (s *recipeScraper) logOnScraped() {
	fmt.Printf("\n\n\n============Recipe============\n\n\n")
	fmt.Printf("URL: %s\n\n", s.recipe.URL)

	fmt.Printf("Title: %s\n\n", s.recipe.Title)
	fmt.Printf("Directions\n\n")
	for i := 0; i < len(s.recipe.Directions); i++ {
		fmt.Printf("Step %v: %s\n", i+1, s.recipe.Directions[i])
	}
	fmt.Printf("Ingredients\n\n")
	for i := 0; i < len(s.recipe.Ingredients); i++ {
		fmt.Printf("%s\n", s.recipe.Ingredients[i])
	}
}
