package jsonapilambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gonobo/jsonapi"
)

// Adapter allows you to use a gibbon jsonapi handler as a lambda handler.
// It leverages the aws-lambda-go-api-proxy library to handle incoming
// API Gateway requests and forwards them to the provided jsonapi handler.
type Adapter struct {
	adapter *httpadapter.HandlerAdapter
}

// Handler handles incoming API Gateway requests. This method should be passed into the lambda.Start()
// function from the github.com/aws/aws-lambda-go/lambda package.
func (m Adapter) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return m.adapter.ProxyWithContext(ctx, request)
}

// Options functions configure the Adapter.
type Options = func(*Adapter)

// New creates a new Adapter with the provided jsonapi handler.
func NewAdapter(handler jsonapi.RequestHandler, options ...func(*jsonapi.H)) Adapter {
	adapter := Adapter{
		adapter: httpadapter.New(jsonapi.Handler(handler, options...)),
	}
	return adapter
}
