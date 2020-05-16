package main

import (
	"context"
	"frank_server/api"
	"frank_server/cache/dynamo"
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

// It is a best practice to instantiate the dynamoDB client outside of the lambda function handler.
// https://aws.amazon.com/blogs/database/building-enterprise-applications-using-amazon-dynamodb-aws-lambda-and-golang/
var cacheStore = dynamo.NewDynamoStore(env)

func main() {
	lambda.Start(handleRequest)
}

// TODO: Move to handler
func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// route the requests
	if strings.Contains(req.Path, "feelingHungry") {
		return handleFeelingHungry()
	}

	// TODO: Move this into separate function
	recipeName := req.QueryStringParameters[recipeParam]
	recipeCount, _ := strconv.Atoi(req.QueryStringParameters[countParam])
	// Set default max search recipes
	if recipeCount > defaultRecipeCount || recipeCount <= 0 {
		recipeCount = defaultRecipeCount
	}

	// TODO: better match search params with rest api so we don't have to modify FE code
	// when local dev vs prod dev
	recipeName = strings.ToLower(recipeName)

	fetcher := api.NewRecipeFetcherService(cacheStore, recipeName, recipeCount)
	recipesView := fetcher.Run()
	recipesViewJSON, err := recipesView.ToJSONString()
	if err != nil {
		log.Fatal(err)
	}
	return buildAPIGatewayProxyResponse(recipesViewJSON), nil
}

func buildAPIGatewayProxyResponse(recipesViewJSON string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       recipesViewJSON,
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
