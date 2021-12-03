package main

import (
	"context"
	"frank_server/cache/dynamo"
	"frank_server/models"
	"frank_server/runner"
	"frank_server/runner/allrecipes"
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
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
	recipeCount, err := strconv.Atoi(req.QueryStringParameters[countParam])
	if err != nil {
		log.Panic(err)
	}

	runner := runner.NewSearchRunner(
		recipeName,
		recipeCount,
		cacheStore,
		allrecipes.NewAllRecipesScraper(colly.NewCollector(), allrecipes.DefaultBuildSearchUrl))

	recipes := runner.Run()

	return buildAPIGatewayProxyResponse(recipes), nil
}

func buildAPIGatewayProxyResponse(recipes models.RecipesView) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string(recipes.Marshal()),
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
