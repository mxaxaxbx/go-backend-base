package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mxaxaxbx/go-backend-base/models"
	"github.com/mxaxaxbx/go-backend-base/server"
	"github.com/mxaxaxbx/go-backend-base/utils"
)

func CheckAuth(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&models.MainResponse{
					Code:    http.StatusUnauthorized,
					Message: "Invalid token",
					Data:    make(map[string]string),
				})
				return
			}

			str := strings.Split(authorization, " ")
			tokenstr := strings.TrimSpace(str[1])
			JWTSecret := os.Getenv("JWTSECRET")

			_, err := jwt.ParseWithClaims(
				tokenstr,
				&models.AppClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(JWTSecret), nil
				},
			)

			if err != nil {
				utils.Log(err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&models.MainResponse{
					Code:    http.StatusUnauthorized,
					Message: "Invalid token",
					Data:    make(map[string]string),
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
