package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/example/auth"
	"github.com/example/utils"
)

// AuthMiddleware checks for a valid JWT in the "Authorization" header
// and adds the user's email to the context if the JWT is valid.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get JWT from "Authorization" header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Split header into "Bearer" and token parts
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.ErrorResponse(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// Parse JWT claims
		claims := &auth.Claims{}
		token, err := auth.ParseToken(tokenString, claims)
		if err != nil || !token.Valid {
			utils.ErrorResponse(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user email to context
		ctx := context.WithValue(r.Context(), "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
