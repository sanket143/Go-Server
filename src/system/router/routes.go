package router

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sanket143/lisa/pkg/types/routes"
	HomeHandler "github.com/sanket143/lisa/src/controllers/home"
)

// Middleware function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes(db *dynamodb.DynamoDB) routes.Routes {
	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
