package allrecipes

import (
	"fmt"
	"frank_server/models"
	"frank_server/runner"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

const baseUrl = "https://www.allrecipes.com"
const searchPath = "/search/results/?search="

func DefaultBuildSearchUrl(recipeName string) string {
	return baseUrl + searchPath + formatUrlSpaces(recipeName)
}

type ArScraper struct {
	Collector      *colly.Collector
	BuildSearchUrl func(string) string
}

func NewAllRecipesScraper(c *colly.Collector, buildSearchUrl func(string) string) runner.Scraper {
	return &ArScraper{Collector: c, BuildSearchUrl: buildSearchUrl}
}

func (s *ArScraper) GetLinks(url string, recipeCount int) map[string]struct{} {

	// use a map here to avoid duplicates
	result := make(map[string]struct{})
	i := 0

	s.Collector.OnHTML(".card__titleLink", func(e *colly.HTMLElement) {
		if i > recipeCount {
			return
		}

		result[e.Attr("href")] = struct{}{}
		i++
	})

	s.Collector.Visit(url)

	s.Collector.Wait()
	return result
}

func (s *ArScraper) GetRecipes(recipeName string, recipeCount int) []*models.Recipe {
	links := s.GetLinks(s.BuildSearchUrl(recipeName), recipeCount)

	result := make([]*models.Recipe, recipeCount)
	i := 0
	// TODO: run this concurrently
	for url := range links {
		recipe := &models.Recipe{}
		s.Collector.OnHTML("li, h1, span", func(e *colly.HTMLElement) {
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
		})

		s.Collector.OnError(func(_ *colly.Response, err error) {
			log.Println("Error scraping:", err)
		})

		s.Collector.OnScraped(func(r *colly.Response) {
			fmt.Printf("scraped\n")
		})

		// Start scraping
		s.Collector.Visit(url)
		s.Collector.Wait()
		result[i] = recipe
		i++
	}

	return result
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
	return strings.Replace(s, " ", "%20", -1)
}

func normalizeString(s string) string {
	return removeExtraSpacesAndNewlines(removeUnneededWords(s))
}

func removeUnneededWords(s string) string {
	// TODO: make this a list of words
	temp := strings.Replace(s, "Advertisement", "", -1)
	return strings.Replace(temp, "Add all ingredients to list", "", -1)
}

func removeExtraSpacesAndNewlines(s string) string {
	return strings.Replace(strings.Join(strings.Fields(s), " "), "\n", "", -1)
}