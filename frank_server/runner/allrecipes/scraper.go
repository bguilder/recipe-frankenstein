package allrecipes

import (
	"fmt"
	"frank_server/models"
	"frank_server/runner"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

const baseURL = "https://www.allrecipes.com"
const searchPath = "/search/results/?search="

// Scraper is an implementation of a runner.Scraper for the All Recipes website.
type Scraper struct {
	BuildSearchURL func(string) string
	Transport      http.RoundTripper
	mu             sync.Mutex
}

// NewAllRecipesScraper returns a new default instance of the AllRecipesScraper
func NewAllRecipesScraper() runner.Scraper {
	return &Scraper{BuildSearchURL: defaultBuildSearchURL, Transport: http.DefaultTransport}
}

func (s *Scraper) GetLinks(url string, recipeCount int) map[string]struct{} {

	c := s.newCollector()

	// use a map here to avoid duplicates
	result := make(map[string]struct{})
	i := 0

	c.OnHTML(".card__titleLink", func(e *colly.HTMLElement) {
		if i > recipeCount {
			return
		}

		result[e.Attr("href")] = struct{}{}
		i++
	})

	if err := c.Visit(url); err != nil {
		panic(err)
	}

	c.Wait()

	return result
}

func (s *Scraper) GetRecipes(recipeName string, recipeCount int) []*models.Recipe {
	links := s.GetLinks(s.BuildSearchURL(recipeName), recipeCount)

	result := []*models.Recipe{}
	wg := sync.WaitGroup{}

	for url := range links {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			recipe := s.scrapeRecipePage(url)

			s.mu.Lock()
			result = append(result, recipe)
			s.mu.Unlock()
		}(url)
	}

	wg.Wait()

	return result
}

func defaultBuildSearchURL(recipeName string) string {
	return baseURL + searchPath + formatUrlSpaces(recipeName)
}

func (s *Scraper) newCollector() *colly.Collector {
	c := colly.NewCollector()
	c.WithTransport(s.Transport)

	return c
}

func (s *Scraper) scrapeRecipePage(url string) *models.Recipe {
	c := s.newCollector()

	recipe := &models.Recipe{}

	c.OnHTML("li, h1, span", func(e *colly.HTMLElement) {
		updateRecipeOnHtml(recipe, e)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error scraping:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("scraped\n")
	})

	// Start scraping
	if err := c.Visit(url); err != nil {
		panic(err)
	}

	c.Wait()

	return recipe
}

func updateRecipeOnHtml(recipe *models.Recipe, e *colly.HTMLElement) {
	if title := tryGetTitle(e); title != "" {
		recipe.Title = title
		return
	}

	if ingredient := tryGetIngredient(e); ingredient != "" {
		recipe.AppendIngredient(ingredient)
		return
	}

	if direction := tryGetDirection(e); direction != "" {
		recipe.AppendDirection(direction)
		return
	}
}

// tryGetIngredient satisfies the source.IRecipe interface
func tryGetIngredient(e *colly.HTMLElement) string {
	if e.DOM.HasClass("ingredients-item-name") {
		return normalizeString(e.Text)
	} else if e.DOM.HasClass("checkList__line") {
		return normalizeString(e.Text)
	}

	return ""
}

// tryGetTitle satisfies the source.IRecipe interface
func tryGetTitle(e *colly.HTMLElement) string {
	if e.Name == "h1" {
		return normalizeString(e.Text)
	}

	return ""
}

// tryGetDirection satisfies the source.IRecipe interface
func tryGetDirection(e *colly.HTMLElement) string {
	if e.DOM.HasClass("recipe-directions__list--item") {
		return normalizeString(e.Text)
	} else if e.DOM.HasClass("instructions-section-item") {
		return normalizeString(e.Text)
	}

	return ""
}

func formatUrlSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func normalizeString(s string) string {
	return removeExtraSpacesAndNewlines(removeUnneededWords(s))
}

func removeUnneededWords(s string) string {
	// TODO: make this a list of words
	temp := strings.ReplaceAll(s, "Advertisement", "")
	return strings.ReplaceAll(temp, "Add all ingredients to list", "")
}

func removeExtraSpacesAndNewlines(s string) string {
	return strings.ReplaceAll(strings.Join(strings.Fields(s), " "), "\n", "")
}
