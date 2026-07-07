package middleware

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"net/http"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/juraevibrahim01/jura/internal/models"
// 	"github.com/juraevibrahim01/jura/internal/service"
// )

// var response models.Auth_middleware_res

// const ClaimsKey models.Contextkey = "claims"

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Auth")

// 		if authHeader == "" {
// 			w.WriteHeader(400)
// 			response.Status = "error"
// 			response.Description = "Auth empty"

// 			//	response
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		claims, err := service.ValidateToken(authHeader, "")
// 		if err != nil {
// 			response.Status = "error"

// 			if errors.Is(err, jwt.ErrTokenExpired) {
// 				response.Description = models.ErrTokenExpired.Error()
// 				w.WriteHeader(400)
// 			} else if err.Error() == models.ErrTokenInvalid.Error() {
// 				response.Description = models.ErrTokenInvalid.Error()
// 				w.WriteHeader(500)
// 			}
// 			// response
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		context := context.WithValue(
// 			r.Context(),
// 			ClaimsKey,
// 			claims,
// 		)

// 		next.ServeHTTP(w, r.WithContext(context))
// 	})
// }
