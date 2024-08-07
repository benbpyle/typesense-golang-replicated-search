package lib

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
)

// CreateUpdateRecipe takes input and sends it to DynamoDB for saving
func CreateUpdateRecipe(ctx context.Context, client *dynamodb.DynamoDB, recipe *Recipe) error {
	marshalledEvent, err := dynamodbattribute.MarshalMap(recipe)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      marshalledEvent,
		TableName: aws.String("TypsenseDemo"),
	}

	logrus.WithFields(logrus.Fields{
		"input": input,
	}).Debug("Pre save")

	_, err = client.PutItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
