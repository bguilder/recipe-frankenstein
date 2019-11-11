package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

const mainSelector = "li, h1, span"
const domain = "https://www.allrecipes.com"
const path = "/recipe/223042/chicken-parmesan"

// Recipe comment
type Recipe struct {
	Title       string
	Ingredients []string
	Directions  []string
}

// Scraper comment
type Scraper interface {
	Scrape() Recipe
}

// Tool comment
type Tool struct {
	Recipe Recipe
}

// Scrape comment
func (s Tool) Scrape() {
	c := colly.NewCollector()

	c.OnHTML(mainSelector, func(e *colly.HTMLElement) {

		// Get Ingredients
		if e.DOM.HasClass("ingredients-item-name") {
			s.Recipe.Ingredients = append(s.Recipe.Ingredients, normalizeString(e.Text))
		} else if e.DOM.HasClass("checkList__line") {
			s.Recipe.Ingredients = append(s.Recipe.Ingredients, normalizeString(e.Text))
		}

		// Get Title
		if e.Name == "h1" && s.Recipe.Title == "" {
			s.Recipe.Title = e.Text
		}

		// Get Directions
		if e.DOM.HasClass("recipe-directions__list--item") {
			s.Recipe.Directions = append(s.Recipe.Directions, normalizeString(e.Text))
		} else if e.DOM.HasClass("instructions-section-item") {
			s.Recipe.Directions = append(s.Recipe.Directions, normalizeString(e.Text))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Finished!")
		fmt.Printf("Title: %s\n\n", s.Recipe.Title)
		for i := 0; i < len(s.Recipe.Directions); i++ {
			fmt.Printf("%s\n", s.Recipe.Directions[i])
		}
		for i := 0; i < len(s.Recipe.Ingredients); i++ {
			fmt.Printf("%s\n", s.Recipe.Ingredients[i])
		}
	})

	// Start scraping
	c.Visit(domain + path)
}

func normalizeString(s string) string {
	return removeExtraSpacesAndNewlines(removeUnneededWords(s))
}

func removeExtraSpacesAndNewlines(s string) string {
	return strings.Replace(strings.Join(strings.Fields(s), " "), "\n", "", -1)
}

func removeUnneededWords(s string) string {
	return strings.Replace(s, "Advertisement", "", -1)
}
