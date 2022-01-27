package main

import (
	resthandler "frank_server/handlers/rest"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080", "localhost:8080"})),
	)

	router.Handle("/search", http.HandlerFunc(resthandler.Search))
	router.Handle("/feelingHungry", http.HandlerFunc(resthandler.FeelingHungry))
	return router
}

// set up server
func serve(router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8088",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting http server: %s", err.Error())
	}
}
