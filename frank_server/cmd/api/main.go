package main

import (
	"encoding/json"
	"fmt"
	"frank_server/postprocessor"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//router := mux.NewRouter()
	ingredients := openIngredients("../../ingredients_fixtures/empanada.json")
	sanitizer := postprocessor.NewSanitizer()
	p := postprocessor.NewPostProcessor(sanitizer)

	router := mux.NewRouter()

	router.Use(handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:8081"})))
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ingredients := p.Run(ingredients)
		payload, err := json.Marshal(ingredients)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}))
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

func openIngredients(file string) []string {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
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
