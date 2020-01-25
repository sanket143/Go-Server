package home

import "github.com/aws/aws-sdk-go/service/dynamodb"

var db *dynamodb.DynamoDB

// Init variables
func Init(DB *dynamodb.DynamoDB) {
	db = DB
}
