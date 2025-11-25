package utils

import "net/http"

// Middleware is a function that wraps http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middlewares := range middlewares {
		handler = middlewares(handler)
	}
	return handler
}
