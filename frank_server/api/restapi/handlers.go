package restapi

import (
	"encoding/json"
	"fmt"
	"frank_server/api"
	"frank_server/cache/dynamo"
	"frank_server/postprocessor"
	"frank_server/scraper"
	"frank_server/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleFeelingHungry(w http.ResponseWriter, r *http.Request) {
	recipes := utils.OpenIngredients("../../ingredients_fixtures/feeling_hungry.json")
	payload, _ := json.Marshal(recipes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {

	//TODO: Move this into separate function
	q := r.URL.Query()
	recipeName := q.Get("recipe")
	recipeCount, _ := strconv.Atoi(q.Get("count"))
	fmt.Printf("recipeName!!!!! %s", recipeName)
	fmt.Printf("recipeCount!!!!! %v", recipeCount)

	// TODO: better sanitize search params
	recipeName = strings.ToLower(recipeName)

	// TODO: move this higher up so we don't instantiate this each time
	cacheStore := dynamo.NewDynamoStore("test")

	fetcher := api.NewRecipeFetcherService(cacheStore, recipeName, recipeCount)
	recipesView := fetcher.Run()
	writePayload(w, recipesView)
}

func writePayload(w http.ResponseWriter, recipesView api.RecipesView) {
	bytes, err := recipesView.ToJSONBytes()
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func formatIngredients(recipes []*scraper.Recipe) postprocessor.IngredientFrequencyList {
	pp := postprocessor.NewPostProcessor()

	ingredients := []string{}
	for _, recipe := range recipes {
		for _, ing := range recipe.Ingredients {
			ingredients = append(ingredients, ing)
		}
	}
	log.Printf("running pp")

	return pp.Run(ingredients)

}
