package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

var response models.Auth_middleware_res

const ClaimsKey models.Contextkey = "claims"

type authRequestBody struct {
	Email string `json:"email"`
}

func AuthMiddleware(authService *service.Auth_service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Auth")

		var reqBody authRequestBody
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response.Status = "error"
				response.Description = "Invalid request body"
				_ = json.NewEncoder(w).Encode(response)
				return
			}
			if len(bodyBytes) > 0 {
				err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&reqBody)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					response.Status = "error"
					response.Description = "Invalid request body"
					_ = json.NewEncoder(w).Encode(response)
					return
				}
			}
			reqBody.Email = strings.ToLower(strings.TrimSpace(reqBody.Email))
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		if authHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			response.Status = "error"
			response.Description = "Auth empty"
			_ = json.NewEncoder(w).Encode(response)
			return
		}

		claims, err := authService.ValidateToken(authHeader, "")
		if err != nil {
			response.Status = "error"

			if errors.Is(err, jwt.ErrTokenExpired) {
				response.Description = models.ErrTokenExpired.Error()
				w.WriteHeader(http.StatusBadRequest)
			} else if err.Error() == models.ErrTokenInvalid.Error() {
				response.Description = models.ErrTokenInvalid.Error()
				w.WriteHeader(http.StatusInternalServerError)
			}
			_ = json.NewEncoder(w).Encode(response)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			ClaimsKey,
			claims,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoleMiddleware(userService *service.User_service, requiredRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(ClaimsKey).(*models.Claims)
		if !ok || claims == nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":      "error",
				"description": "Unauthorized",
			})
			return
		}

		userRole, err := userService.GetUserRole(claims.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":      "error",
				"description": "Ошибка проверки роли",
			})
			return
		}

		if userRole == "" {
			userRole = "reading"
		}

		userRole = strings.ToLower(strings.TrimSpace(userRole))
		allowed := false
		for _, role := range requiredRoles {
			if strings.ToLower(strings.TrimSpace(role)) == userRole {
				allowed = true
				break
			}
		}

		if !allowed {
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":      "error",
				"description": "Forbidden",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
