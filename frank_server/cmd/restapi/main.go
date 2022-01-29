package main

import (
	"frank_server/db/dynamo"
	handler "frank_server/handler/http"
	"frank_server/runner"
	"frank_server/runner/allrecipes"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//  ---------------------- NOTE -------------------------
// This http server is only used for local dev currently.
// ------------------------------------------------------

const (
	env       = "local"
	serverUrl = "0.0.0.0:8088"
	clientUrl = "http://localhost:8080"
)

func main() {
	router := newRouter()
	serve(router)
}

// set up router
func newRouter() *mux.Router {
	router := mux.NewRouter()

	// Allow CORS for specific origins
	router.Use(handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedOrigins([]string{clientUrl})),
	)

	searchRunner := runner.NewSearchRunner(
		dynamo.NewDynamoStore(env),
		allrecipes.NewAllRecipesScraper(colly.NewCollector(), allrecipes.DefaultBuildSearchUrl))

	router.Handle("/search", http.HandlerFunc(handler.Search(searchRunner)))
	router.Handle("/feelingHungry", http.HandlerFunc(handler.FeelingHungry))
	return router
}

// set up server
func serve(router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         serverUrl,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting http server: %s", err.Error())
	}
}
