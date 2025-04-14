package role_middleware

import (
	"avito-spring-test/internal/models/dto"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const op = "role_middleware.RequireRole"

			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, dto.ErrorResponse{Message: "error authenticating: no token"})
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("%s: invalid token", op)
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, dto.ErrorResponse{Message: "error authenticating: err parse token"})
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				role, err := extractRoleFromClaims(claims)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					render.JSON(w, r, dto.ErrorResponse{Message: "error authenticating: err parse token"})
				}
				for _, allowed := range allowedRoles {
					if role == allowed {
						next.ServeHTTP(w, r)
						return
					}
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, dto.ErrorResponse{Message: "error authenticating: err parse token"})
			}

			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		})
	}
}

func extractRoleFromClaims(claims jwt.MapClaims) (string, error) {
	// Проверяем, что поле userID существует
	roleValue, ok := claims["role"]
	if !ok {
		return "", errors.New("role not found in token")
	}

	// Проверяем тип userID
	switch v := roleValue.(type) {
	case string:
		return v, nil
	default:
		return "", errors.New("role is not a number")
	}
}
