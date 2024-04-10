package auth

import (
	"context"
	"log"
	"net/http"
)

var userIDCtxKey = &contextKey{"userID"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userID := r.Header.Get("X-User-Id")

			if userID == "" {
				log.Println("X-User-Id header not found")
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userIDCtxKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userIDCtxKey).(string)
	println(raw)
	return raw
}
