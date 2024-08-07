package main

import (
	"context"
	"encoding/json"
	"os"
	"typesense-demo/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"github.com/typesense/typesense-go/v2/typesense"
)

var client *typesense.Client

// handler runs with each API Gateway Request
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	statusCode := 200

	query := ""

	// parse out the query string for the search or set it to ""
	if q, ok := request.QueryStringParameters["q"]; ok {
		query = q
	}

	// fetch all documents from Typesense
	r, err := lib.SearchDocuments(ctx, client, query)
	// if an error is returned from the search, log it and return 500
	if err != nil {
		logrus.Errorf("Error searching documents: %s", err)
		statusCode = 500
		return events.APIGatewayProxyResponse{
			StatusCode:        statusCode,
			Headers:           map[string]string{},
			MultiValueHeaders: map[string][]string{},
			Body:              "",
			IsBase64Encoded:   false,
		}, nil
	}

	// conver the documents to DTO structs, marshal them into a JSON string
	// and then return out to the client
	dtos := lib.NewRecipeViewsFromRecipes(r)
	v, _ := json.Marshal(dtos)
	return events.APIGatewayProxyResponse{
		StatusCode:        statusCode,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              string(v),
		IsBase64Encoded:   false,
	}, nil
}

// main is the entry point for the Lambda Function
func main() {
	lambda.Start(handler)
}

// init runs before main and only once
func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	url := os.Getenv("TYPESENSE_CLUSTER_URL")
	apiKey := os.Getenv("TYPESENSE_API_KEY")
	client = typesense.NewClient(
		typesense.WithServer(url),
		typesense.WithAPIKey(apiKey))
}
