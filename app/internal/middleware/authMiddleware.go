package middleware

import (
	"app/internal/common"
	"context"
	"net/http"
)	

// authMiddleware is a middleware that checks for the JWT token in the Authorization header
// It verifies the token and extracts userId and role from the claims.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name,role, err := common.ExtractUserFromJWT(w,r)
		if err != nil {
			common.HandleErrorSimple(w, err, http.StatusUnauthorized)
			return
		}
		
		ctx := r.Context()
		ctx = context.WithValue(ctx,"userId", name)
		ctx = context.WithValue(ctx,"Role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

