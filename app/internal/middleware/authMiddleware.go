package middleware

import (
	"app/config"
	"app/internal/common"
	"context"
	"net/http"
)

type AuthenticateMiddleware struct {
	appConfig *config.AppConfig
}

func NewAutenticateMiddleware(appConfig *config.AppConfig) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{appConfig: appConfig}
}

// AuthMiddleware authMiddleware is a middleware that checks for the JWT token in the Authorization header
// It verifies the token and extracts userId and role from the claims.
func (m *AuthenticateMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name, role, err := common.ExtractUserFromJWT(m.appConfig.JWT.Key, r)
		if err != nil {
			common.HandleErrorSimple(w, err, http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", name)
		ctx = context.WithValue(ctx, "Role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
