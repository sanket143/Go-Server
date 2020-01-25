package router

import (
	"net/http"

	"github.com/sanket143/lisa/pkg/types/routes"
	StatusHandler "github.com/sanket143/lisa/src/controllers/v1/status"
)

// Middleware Function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {

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
