package router

import (
	"github.com/gorilla/mux"
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
			Handler(route.HandleFunc)
	}
}

// NewRouter Creates new Router
func NewRouter() (r Router) {
	r.Router = mux.NewRouter()

	return
}
