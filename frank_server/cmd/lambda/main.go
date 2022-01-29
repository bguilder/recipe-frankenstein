package main

import (
	"context"
	"fmt"
	lambdahandler "frank_server/handler/lambda"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	feelingHungeryRoute = "feelingHungry"
	searchRoute         = "search"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.Contains(req.Path, feelingHungeryRoute) {
		return lambdahandler.HandleFeelingHungry()
	}

	if strings.Contains(req.Path, searchRoute) {
		return lambdahandler.HandleSearch(req)
	}

	return events.APIGatewayProxyResponse{StatusCode: 500}, fmt.Errorf("unknown lambda request with path: %s", req.Path)
}
