package resthandler

import (
	"encoding/json"
	"fmt"
	"frank_server/cache/dynamo"
	"frank_server/runner"
	"frank_server/runner/allrecipes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func FeelingHungry(w http.ResponseWriter, r *http.Request) {
	recipes := OpenIngredients("./ingredients_fixtures/feeling_hungry.json")
	payload, err := json.Marshal(recipes)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	recipeName := q.Get("recipe")
	count := q.Get("count")

	recipeCount, err := strconv.Atoi(count)
	if err != nil {
		log.Panicf("err getting count: %s", err.Error())
	}

	runner := runner.NewSearchRunner(
		recipeName,
		recipeCount,
		dynamo.NewDynamoStore("test"),
		allrecipes.NewAllRecipesScraper(colly.NewCollector(), allrecipes.DefaultBuildSearchUrl))

	recipesView := runner.Run()

	w.Header().Set("Content-Type", "application/json")
	w.Write(recipesView.Marshal())
}

// opens a static file of ingredients
// ../../ingredients_fixtures/empanada.json
func OpenIngredients(file string) []string {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []string
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func serveHttpInternalServerError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
