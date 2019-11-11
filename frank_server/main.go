package main

import (
	"fmt"
	"frank_server/scraper"
	"net/http"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/test", recipeHandler)
	// http.Handle("/", r)

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// srv.ListenAndServe()
	recipe := scraper.Recipe{}
	firstScraper := scraper.Tool{Recipe: recipe}
	firstScraper.Scrape()
}

func recipeHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("test?")
	recipe := scraper.Recipe{}
	firstScraper := scraper.Tool{Recipe: recipe}
	firstScraper.Scrape()
}
