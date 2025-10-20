package middleware

import (
	"context"
	"net/http"
	"strings"
	"bookmark-api/utils"
)

type ctxKeyUserID struct{}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenStr := tokenParts[1]
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

        ctx := r.Context()
        ctx = context.WithValue(ctx, ctxKeyUserID{}, claims.UserID)
		// check if you can do
		// next(w, r.WithContext(ctx))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserIDFromContext(r *http.Request) (uint, bool) {
    v := r.Context().Value(ctxKeyUserID{})
    if v == nil {
        return 0, false
    }
    id, ok := v.(uint)
    return id, ok
}
