package main

import (
	"context"
	"encoding/json"
	"frank_server/models"
	"frank_server/postprocessor"
	"frank_server/runner"
	"frank_server/scraper/allrecipes"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	defaultRecipeCount = 3

	recipeParam              = "recipe"
	countParam               = "count"
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
	recipeFrankDomain        = "http://recipefrankenstein.com"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	recipeName := req.QueryStringParameters[recipeParam]
	recipeCount, _ := strconv.Atoi(req.QueryStringParameters[countParam])
	// Set default max search recipes to 5
	if recipeCount > defaultRecipeCount || recipeCount <= 0 {
		recipeCount = defaultRecipeCount
	}

	runner := runner.NewRunner(recipeName, recipeCount, &allrecipes.SearchScraper{}, &allrecipes.RecipeScraper{})
	recipes := runner.Run()

	return events.APIGatewayProxyResponse{Body: formatIngredients(recipes), StatusCode: 200, Headers: defaultHeaders()}, nil
}

func defaultHeaders() map[string]string {
	headers := make(map[string]string)
	headers[accessControlAllowOrigin] = recipeFrankDomain
	return headers
}

func formatIngredients(recipes []*models.Recipe) string {
	pp := postprocessor.NewPostProcessor(postprocessor.NewSanitizer())

	ingredients := []string{}
	for _, recipe := range recipes {
		for _, ing := range recipe.Ingredients {
			ingredients = append(ingredients, ing)
		}
	}

	result := pp.Run(ingredients)
	recipesView := models.RecipesView{Recipes: recipes, Ingredients: result}
	payload, err := json.Marshal(recipesView)
	if err != nil {
		log.Printf("error marshalling recipe view %v", err)
	}
	return string(payload)

}
