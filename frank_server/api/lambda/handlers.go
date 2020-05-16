package lambdahandler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	store Store
	ctx   context.Context
	req   events.APIGatewayProxyRequest
}

func NewLambdaHandler(store Store, ctx context.Context, req events.APIGatewayProxyRequest) LambdaHandler {
	return LambdaHandler{store: store, ctx: ctx, req: req}
}

func (h *LambdaHandler) Handle() (events.APIGatewayProxyResponse, error) {

}
