package main

import (
	"encoding/json"
	"frank_server/models"
	"frank_server/postprocessor"
	"frank_server/runner"
	"frank_server/scraper/allrecipes"
	"frank_server/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const recipeName = "empanada"
const numberOfRecipes = 10

var cachedRecipes []byte

func main() {

	router := newRouter()

	serve(router)
}

// set up router
func newRouter() *mux.Router {
	router := mux.NewRouter()

	// Allow CORS for specific origins
	router.Use(handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"})),
	)

	router.Handle("/example", http.HandlerFunc(handleExampleIngredients))
	router.Handle("/search", http.HandlerFunc(handleSearch))
	return router
}

// set up server
func serve(router *mux.Router) {
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error 1: %v", err)
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	if cachedRecipes != nil {
		log.Printf("Reading from cache: %s", string(cachedRecipes))
		w.Header().Set("Content-Type", "application/json")
		w.Write(cachedRecipes)
		return
	}

	runner := runner.NewRunner(recipeName, numberOfRecipes, &allrecipes.SearchScraper{}, &allrecipes.RecipeScraper{})
	recipes := runner.Run()
	ing := formatIngredients(recipes)
	log.Printf("ran pp: %+v", ing)

	recipesView := models.RecipesView{Recipes: recipes, Ingredients: ing}

	payload, _ := json.Marshal(recipesView)
	cachedRecipes = payload
	log.Printf("setting cache: %s", string(cachedRecipes))
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func formatIngredients(recipes []*models.Recipe) postprocessor.PairList {
	pp := postprocessor.NewPostProcessor(postprocessor.NewSanitizer())

	ingredients := []string{}
	for _, recipe := range recipes {
		for _, ing := range recipe.Ingredients {
			ingredients = append(ingredients, ing)
		}
	}
	log.Printf("running pp")

	return pp.Run(ingredients)

}

// handle func
func handleExampleIngredients(w http.ResponseWriter, r *http.Request) {
	exIngredients := getExampleIngredients()
	w.Header().Set("Content-Type", "application/json")
	w.Write(exIngredients)
}

func getExampleIngredients() []byte {
	ingredients := utils.OpenIngredients("../ingredients_fixtures/empanada.json")
	postProcessor := postprocessor.NewPostProcessor(postprocessor.NewSanitizer())
	formattedIngredients := postProcessor.Run(ingredients)
	payload, err := json.Marshal(formattedIngredients)
	if err != nil {
		log.Println(err)
	}
	return payload
}