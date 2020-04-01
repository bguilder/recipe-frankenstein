package main

import (
	"context"
	"encoding/json"
	"frank_server/cache/dynamo"
	"frank_server/models"
	"frank_server/postprocessor"
	"frank_server/runner"
	"frank_server/scraper"
	"frank_server/source/allrecipes"
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	defaultRecipeCount = 7

	recipeParam              = "recipe"
	countParam               = "count"
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
	recipeFrankDomain        = "http://recipefrankenstein.com"
	env                      = "live"
)

// //It is a best practice to instanciate the dynamoDB client outside of the lambda function handler.
// https://aws.amazon.com/blogs/database/building-enterprise-applications-using-amazon-dynamodb-aws-lambda-and-golang/
var cacheStore = dynamo.NewDynamoStore(env)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// route the requests
	if strings.Contains(req.Path, "feelingHungry") {
		return handleFeelingHungry()
	}

	recipeName := req.QueryStringParameters[recipeParam]
	recipeCount, _ := strconv.Atoi(req.QueryStringParameters[countParam])
	// Set default max search recipes
	if recipeCount > defaultRecipeCount || recipeCount <= 0 {
		recipeCount = defaultRecipeCount
	}

	recipes, err := cacheStore.GetRecipes(recipeName)
	if err != nil {
		log.Panic(err)
	}
	if recipes != nil {
		log.Println("loaded from cache")
		return buildAPIGatewayProxyResponse(recipes), nil
	}
	log.Println("cache miss...")

	runner := runner.NewRunner(
		recipeName,
		recipeCount,
		scraper.NewLinkScraper(&allrecipes.LinkSource{}),
		scraper.NewRecipeScraper(&allrecipes.RecipeSource{}))

	recipes = runner.Run()

	// update cache
	err = cacheStore.PutRecipes(recipeName, recipes)
	if err != nil {
		log.Panic(err)
	}

	return buildAPIGatewayProxyResponse(recipes), nil
}

func buildAPIGatewayProxyResponse(recipes []*scraper.Recipe) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       formatIngredients(recipes),
		StatusCode: 200,
		Headers:    defaultHeaders()}
}

func handleFeelingHungry() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: feelingHungry, StatusCode: 200, Headers: defaultHeaders()}, nil
}

func defaultHeaders() map[string]string {
	headers := make(map[string]string)
	headers[accessControlAllowOrigin] = recipeFrankDomain
	return headers
}

func formatIngredients(recipes []*scraper.Recipe) string {
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

// TODO: Move to static file
const feelingHungry = `[
	"Chicken Parmesan",
	"Fish Tacos",
	"Chicken and Dumplings",
	"Coq Au Vin",
	"Jambalaya",
	"Empanadas",
	"Stir Fry",
	"Burritos",
	"Fajitas",
	"Burger",
	"Quesadillas",
	"KFC Bowl",
	"Sloppy Joes",
	"Chicken and Rice",
	"Pulled BBQ Chicken",
	"Pot Roast",
	"Ribs",
	"French dip sandwiches",
	"Pizza",
	"Chicken Pot Pie",
	"Shepherd's Pie",
	"Calzones",
	"Enchiladas",
	"Stuffed Green Peppers",
	"Chicken Cordon Bleu",
	"Breaded chicken",
	"Chicken wings",
	"Lemon Chicken",
	"Pork Chops",
	"Chimichangas",
	"Meatloaf",
	"Meatball Subs",
	"Stuffed Pork Chops",
	"Beef & Broccoli",
	"Ham and Cabbage",
	"Salmon",
	"Fish Sticks",
	"Fish ‘N Chips",
	"Crab Cakes",
	"Mac ‘N Cheese",
	"Lasagna",
	"Baked Ziti",
	"Fettuccine Alfredo",
	"Manicotti",
	"Spaghetti and Meatballs",
	"Pesto Pasta",
	"Eggplant Parmesan",
	"Tortellini",
	"Stroganoff",
	"Pad Tai",
	"Minestrone Soup",
	"Chicken Noodle Soup",
	"Clam Chowder",
	"Tomato Bisque",
	"Italian Wedding Soup",
	"French Onion Soup",
	"Panini",
	"BBQ Chicken Salad",
	"Pasta Salad",
	"Macaroni Salad",
	"Pulled Chicken",
	"Potato Salad",
	"Rib-eye Steak",
	"Strip Steak",
	"Kabobs",
	"French Toast",
	"Waffles ",
	"Crepes",
	"Omelet"
]`
