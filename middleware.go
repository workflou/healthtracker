package main

import "net/http"

type middlewareFunc func(http.Handler) http.Handler

func newAuthMiddleware() middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		})
	}
}
