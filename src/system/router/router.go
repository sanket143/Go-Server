package router

import (
	"github.com/gorilla/mux"
	"github.com/sanket143/lisa/pkg/types/routes"
	apiSubRoutes "github.com/sanket143/lisa/src/controllers/api/router"
)

// Router Model
type Router struct {
	Router *mux.Router
}

// Init all handlers
func (r *Router) Init() {
	r.Router.Use(Middleware)

	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	apiSubRoutes := apiSubRoutes.GetRoutes()
	for name, pack := range apiSubRoutes {
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
