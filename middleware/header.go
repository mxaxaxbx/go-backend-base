package middleware

import (
	"context"
	"net/http"

	"github.com/mxaxaxbx/go-backend-base/server"
)

func AddHeaders(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "server", s)

			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
