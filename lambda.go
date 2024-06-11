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
	adapter   *httpadapter.HandlerAdapter
	adapterV2 *httpadapter.HandlerAdapterV2
}

// Handler handles incoming API Gateway requests. This method should be passed into the lambda.Start()
// function from the github.com/aws/aws-lambda-go/lambda package.
func (a Adapter) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return a.adapter.ProxyWithContext(ctx, request)
}

// HandlerV2 handles incoming API Gateway V2 requests. This method should be passed into the lambda.Start()
// function from the github.com/aws/aws-lambda-go/lambda package.
func (a Adapter) HandlerV2(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return a.adapterV2.ProxyWithContext(ctx, request)
}

// Options functions configure the Adapter.
type Options = func(*Adapter)

// NewAdapter creates a new Adapter with the provided jsonapi handler.
func NewAdapter(handler jsonapi.RequestHandler, options ...func(*jsonapi.H)) Adapter {
	adapter := Adapter{
		adapter: httpadapter.New(jsonapi.Handler(handler, options...)),
	}
	return adapter
}

// NewAdapterV2 creates a new Adapter with the provided jsonapi handler.
func NewAdapterV2(handler jsonapi.RequestHandler, options ...func(*jsonapi.H)) Adapter {
	adapter := Adapter{
		adapterV2: httpadapter.NewV2(jsonapi.Handler(handler, options...)),
	}
	return adapter
}
