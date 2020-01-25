package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Connect to DynamoDB
func Connect() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	dynamoConfig := &aws.Config{Endpoint: aws.String("http://localhost:8000")}
	svc := dynamodb.New(sess, dynamoConfig)

	return svc
}
