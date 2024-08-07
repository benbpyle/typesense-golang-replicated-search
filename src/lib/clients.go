package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewDynamoDBClient inits a DynamoDB session to be used throughout the services
func NewDynamoDBClient() *dynamodb.DynamoDB {
	c := &aws.Config{
		Region: aws.String("us-west-2"),
	}

	sess := session.Must(session.NewSession(c))

	return dynamodb.New(sess)
}
