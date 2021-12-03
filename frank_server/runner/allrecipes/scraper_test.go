package allrecipes_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"frank_server/runner/allrecipes"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
)

func TestScraper_GetLinks(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(transport)

	s := allrecipes.ArScraper{Collector: c}
	fmt.Println(dir)
	links := s.GetLinks("file://"+dir+"/html/chicken_marsala_links.html", 2)
	assert.Equal(t, 2, len(links))
	assert.Contains(t, links, "file:///home/brianguilder/dev/projects/recipe-frankenstein/frank_server/runner/allrecipes/html/child_page/one.html")
}

func buildSearchUrl(dir string) string {
	return "file://" + dir + "/html/chicken_marsala_links.html"
}

func TestScraper_GetRecipes(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(transport)

	s := allrecipes.ArScraper{Collector: c, BuildSearchUrl: buildSearchUrl}
	recipes := s.GetRecipes(dir, 2)
	assert.Equal(t, 2, len(recipes))
	assert.Equal(t, "Chicken Marsala with Portobello Mushrooms", recipes[0].Title)
	assert.Equal(t, 13, len(recipes[0].Ingredients))
	assert.Equal(t, 6, len(recipes[0].Directions))
	assert.Equal(t, "Easy Chicken Marsala", recipes[1].Title)
	assert.Equal(t, 8, len(recipes[1].Ingredients))
	assert.Equal(t, 3, len(recipes[1].Directions))

}
