package main

import (
	"context"
	"os"
	"typesense-demo/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"github.com/typesense/typesense-go/v2/typesense"
)

var client *typesense.Client

//	 handler is the main lambda function that runs when a DynamoDBEvent is triggered
//		via the DDB stream
func handler(ctx context.Context, event events.DynamoDBEvent) (interface{}, error) {
	logrus.WithFields(logrus.Fields{
		"event": event,
	}).Info("The Event")

	for _, v := range event.Records {
		if v.EventName == "REMOVE" {
			continue
		}

		recipe := lib.NewRecipeFromStreamRecord(v)
		logrus.WithFields(logrus.Fields{
			"recipe": recipe,
		}).Info("Recipe made")

		typesenseRecipe := lib.NewRecipeTypesenseFromRecipe(recipe)
		_, err := client.Collection("recipes").Documents().Upsert(ctx, typesenseRecipe)
		if err != nil {
			logrus.Errorf("Error creating new Typesense document: %s", err)
		}
	}
	return nil, nil
}

// main is the entry point for the Lambda Function
func main() {
	lambda.Start(handler)
}

// init runs as the function launches and before more.  Sets up various required elements
func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	url := os.Getenv("TYPESENSE_CLUSTER_URL")
	apiKey := os.Getenv("TYPESENSE_API_KEY")
	client = typesense.NewClient(
		typesense.WithServer(url),
		typesense.WithAPIKey(apiKey))
}
