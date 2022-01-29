package httphandler

import (
	"fmt"
	"frank_server/handler"
	"frank_server/runner"
	"net/http"
	"strconv"
)

func FeelingHungry(w http.ResponseWriter, r *http.Request) {
	result, err := handler.OpenFeelingHungry()
	if err != nil {
		serveHttpInternalServerError(err, w)
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(result); err != nil {
		serveHttpInternalServerError(err, w)
	}
}

func Search(searchRunner runner.SearchRunner) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		recipeName := q.Get("recipe")
		count := q.Get("count")

		recipeCount, err := strconv.Atoi(count)
		if err != nil {
			serveHttpInternalServerError(err, w)
		}

		recipesView := searchRunner.Run(recipeName, recipeCount)

		w.Header().Set("Content-Type", "application/json")
		w.Write(recipesView.Marshal())
	}
}

func serveHttpInternalServerError(err error, w http.ResponseWriter) {
	fmt.Printf("error serving data: %s", err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
