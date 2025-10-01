package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Black-tag/productAPI/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Authenticate(secret string, db *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "no header", http.StatusUnauthorized)
				return
			}
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 {
				http.Error(w, "malfromed troken", http.StatusUnauthorized)
				return
			}
			if parts[0] != "Bearer" {
				http.Error(w, "header must contain bearer", http.StatusUnauthorized)
				return
			}
			tokenSring := parts[1]
			claims := &jwt.RegisteredClaims{}
			token, err := jwt.ParseWithClaims(tokenSring, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				http.Error(w, "error parsing token", http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time) {
				http.Error(w, "token has expired", http.StatusUnauthorized)
				return
			}
			userID, err := uuid.Parse(claims.Subject)
			if err != nil {
				http.Error(w, "error parsing user id", http.StatusInternalServerError)
				return
			}

			role, err := db.GetRoleByID(r.Context(), userID)
			if err != nil {
				http.Error(w, "unable to fetch role", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(ctx, "tokenString", tokenSring)
			ctx = context.WithValue(ctx, "role", role)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
