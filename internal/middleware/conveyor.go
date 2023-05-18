package middleware

import "net/http"

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Conveyor(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
