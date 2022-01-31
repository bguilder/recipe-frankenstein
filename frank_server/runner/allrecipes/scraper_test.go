package allrecipes_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"frank_server/models"
	"frank_server/runner/allrecipes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type scraperSuite struct {
	suite.Suite
}

func TestRunScraperSuite(t *testing.T) {
	suite.Run(t, new(scraperSuite))
}

func (s *scraperSuite) SetupTest() {
	replaceFileLocation()
}

func (s *scraperSuite) TearDownTest() {
	replaceOldFileLocation()
}

func (s *scraperSuite) TestScraper_GetLinks() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	scraper := allrecipes.Scraper{Transport: transport}
	fmt.Println(dir)
	links := scraper.GetLinks(buildSearchURL(dir), 2)
	assert.Equal(s.T(), 2, len(links))
	assert.Contains(s.T(), links, "file://"+dir+"/html/child_page/one.html")
	assert.Contains(s.T(), links, "file://"+dir+"/html/child_page/two.html")
}

func buildSearchURL(dir string) string {
	return "file://" + dir + "/html/chicken_marsala_links.html"
}

func (s *scraperSuite) TestScraper_GetRecipes() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	scraper := allrecipes.Scraper{BuildSearchURL: buildSearchURL, Transport: transport}
	recipes := scraper.GetRecipes(dir, 2)
	assert.Equal(s.T(), 2, len(recipes))
	portobello, easy := findRecipes(recipes)
	assert.Equal(s.T(), "Chicken Marsala with Portobello Mushrooms", portobello.Title)
	assert.Equal(s.T(), 13, len(portobello.Ingredients))
	assert.Equal(s.T(), 6, len(portobello.Directions))
	assert.Equal(s.T(), "Easy Chicken Marsala", easy.Title)
	assert.Equal(s.T(), 8, len(easy.Ingredients))
	assert.Equal(s.T(), 3, len(easy.Directions))
}

// findRecipes returns portobello in first return spot and easy in the second
func findRecipes(recipes []*models.Recipe) (*models.Recipe, *models.Recipe) {
	if strings.Contains(recipes[0].Title, "Portobello") {
		return recipes[0], recipes[1]
	}
	return recipes[1], recipes[0]
}

func replaceFileLocation() {
	// open the file
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file := dir + "/html/chicken_marsala_links.html"

	// find the replacement string key
	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input,
		[]byte("{ReplaceMe}"),
		[]byte(dir), -1)

	if err = ioutil.WriteFile(file, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// put the location of the actual file
}

func replaceOldFileLocation() {
	// open the file
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file := dir + "/html/chicken_marsala_links.html"

	// find the replacement string key
	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input,
		[]byte(dir),
		[]byte("{ReplaceMe}"), -1)

	if err = ioutil.WriteFile(file, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// put the location of the actual file
}
