package routes

import "net/http"

// Routes Array of Route
type Routes []Route

// Route model
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// SubRoutePackage model
type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
