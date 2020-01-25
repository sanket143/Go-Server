package status

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

// Init variables
func Init(DB *dynamodb.DynamoDB) {
	db = DB
}

// Index Function for Status
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Status: active"))
}
