package main

import (
	"encoding/json"
	"fmt"
	"frank_server/cache/dynamo"
	"frank_server/postprocessor"
	"frank_server/runner"
	"frank_server/runner/allrecipes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const numberOfRecipes = 1

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
		handlers.AllowedOrigins([]string{"http://localhost:8080", "localhost:8080"})),
	)
	// TODO: Make these query params to match lambda
	router.Handle("/example", http.HandlerFunc(handleExampleIngredients))
	router.Handle("/search", http.HandlerFunc(handleSearch))
	router.Handle("/feelingHungry", http.HandlerFunc(handleFeelingHungry))
	return router
}

// set up server
func serve(router *mux.Router) {
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error 1: %v", err)
	}
}

func handleFeelingHungry(w http.ResponseWriter, r *http.Request) {
	recipes := OpenIngredients("./ingredients_fixtures/feeling_hungry.json")
	payload, _ := json.Marshal(recipes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
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

// handle func
func handleExampleIngredients(w http.ResponseWriter, r *http.Request) {
	exIngredients := getExampleIngredients()
	w.Header().Set("Content-Type", "application/json")
	w.Write(exIngredients)
}

func getExampleIngredients() []byte {
	ingredients := OpenIngredients("./ingredients_fixtures/empanada.json")
	postProcessor := postprocessor.NewPostProcessor()
	formattedIngredients := postProcessor.Run(ingredients)
	payload, err := json.Marshal(formattedIngredients)
	if err != nil {
		log.Println(err)
	}
	return payload
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
