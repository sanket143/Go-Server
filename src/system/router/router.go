package router

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
	"github.com/sanket143/lisa/pkg/types/routes"
	v1SubRoutes "github.com/sanket143/lisa/src/controllers/v1/router"
)

// Router Model
type Router struct {
	Router *mux.Router
}

// Init all handlers
func (r *Router) Init(db *dynamodb.DynamoDB) {
	r.Router.Use(Middleware)

	baseRoutes := GetRoutes(db)
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	v1SubRoutes := v1SubRoutes.GetRoutes(db)
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

// AttachSubRouterWithMiddleware Function
func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

// NewRouter Creates new Router
func NewRouter() (r Router) {
	r.Router = mux.NewRouter()

	return
}
