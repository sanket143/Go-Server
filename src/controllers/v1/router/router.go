package router

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sanket143/lisa/pkg/types/routes"
	StatusHandler "github.com/sanket143/lisa/src/controllers/v1/status"
)

var db *dynamodb.DynamoDB

// Middleware Function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")

		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")
		log.Println(r.Header.Get("X-App-Token"))

		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes(DB *dynamodb.DynamoDB) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB

	StatusHandler.Init(DB)

	/* Routes */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}

	return
}
