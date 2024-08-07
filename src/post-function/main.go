package main

import (
	"context"
	"encoding/json"
	"typesense-demo/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
)

var client *dynamodb.DynamoDB

// handler is the entry point that runs when the Lambda Function is trigger by an API Gateway Request
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dto := &lib.RecipeCreateDto{}
	json.Unmarshal([]byte(request.Body), dto)
	recipe := lib.NewRecipeFromCreate(*dto)
	statusCode := 200
	err := lib.CreateUpdateRecipe(ctx, client, recipe)
	if err != nil {
		logrus.Errorf("(Error)=%v", err)
		statusCode = 500
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        statusCode,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              "",
		IsBase64Encoded:   false,
	}, nil
}

// main runs when your Lambda is launched and before the handler is executed
func main() {
	lambda.Start(handler)
}

// init runs before main and only once
func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	client = lib.NewDynamoDBClient()
}
