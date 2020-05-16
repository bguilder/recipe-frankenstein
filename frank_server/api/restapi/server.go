package restapi

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Serve starts the http server
func Serve() {
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
		handlers.AllowedOrigins([]string{"http://localhost:8080"})),
	)
	// TODO: Make these query params to match lambda
	router.Handle("/search", http.HandlerFunc(handleSearch))
	router.Handle("/feelingHungry", http.HandlerFunc(handleFeelingHungry))
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
