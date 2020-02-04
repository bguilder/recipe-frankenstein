package main

import (
	"context"
	"encoding/json"
	"frank_server/models"
	"frank_server/postprocessor"
	"frank_server/runner"
	"frank_server/scraper/allrecipes"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const numberOfRecipes = 1

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	recipeName := req.QueryStringParameters["recipe"]
	recipeCount, _ := strconv.Atoi(req.QueryStringParameters["count"])
	// Set max search recipes to 5 for now
	if recipeCount < 5 {
		recipeCount = 5
	}

	runner := runner.NewRunner(recipeName, recipeCount, &allrecipes.SearchScraper{}, &allrecipes.RecipeScraper{})
	recipes := runner.Run()
	ing := formatIngredients(recipes)
	recipesView := models.RecipesView{Recipes: recipes, Ingredients: ing}

	payload, _ := json.Marshal(recipesView)
	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"
	return events.APIGatewayProxyResponse{Body: string(payload), StatusCode: 200, Headers: headers}, nil
}

func formatIngredients(recipes []*models.Recipe) postprocessor.PairList {
	pp := postprocessor.NewPostProcessor(postprocessor.NewSanitizer())

	ingredients := []string{}
	for _, recipe := range recipes {
		for _, ing := range recipe.Ingredients {
			ingredients = append(ingredients, ing)
		}
	}

	return pp.Run(ingredients)

}
