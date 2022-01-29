package lambda

import (
	"frank_server/cache/dynamo"
	"frank_server/handler"
	"frank_server/runner"
	"frank_server/runner/allrecipes"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gocolly/colly"
)

const (
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
	recipeFrankDomain        = "http://recipefrankenstein.com"
	recipeParam              = "recipe"
	countParam               = "count"
	env                      = "live"
)

// //It is a best practice to instanciate the dynamoDB client outside of the lambda function handler.
// https://aws.amazon.com/blogs/database/building-enterprise-applications-using-amazon-dynamodb-aws-lambda-and-golang/
var cacheStore = dynamo.NewDynamoStore(env)

var lambdaResponseHeader = map[string]string{accessControlAllowOrigin: recipeFrankDomain}

func HandleFeelingHungry() (events.APIGatewayProxyResponse, error) {
	result, err := handler.OpenFeelingHungry()
	if err != nil {
		log.Panic(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    lambdaResponseHeader,
		Body:       string(result),
	}, nil
}

func HandleSearch(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	recipeName := req.QueryStringParameters[recipeParam]
	recipeCount, err := strconv.Atoi(req.QueryStringParameters[countParam])
	if err != nil {
		log.Panic(err)
	}

	runner := runner.NewSearchRunner(
		cacheStore,
		allrecipes.NewAllRecipesScraper(colly.NewCollector(), allrecipes.DefaultBuildSearchUrl))

	recipes := runner.Run(recipeName, recipeCount)

	return events.APIGatewayProxyResponse{
		Body:       string(recipes.Marshal()),
		StatusCode: 200,
		Headers:    lambdaResponseHeader}, nil
}
