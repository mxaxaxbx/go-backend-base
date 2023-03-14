package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mxaxaxbx/go-backend-base/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	claims := models.AppClaims{
		UserId: "",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
		},
	}

	JWTSecret := os.Getenv("JWTSECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.MainResponse{
			Code:    http.StatusInternalServerError,
			Message: "An error ocurred. Try again later",
			Data:    make(map[string]string),
		})
		return
	}

	json.NewEncoder(w).Encode(&models.MainResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    tokenString,
	})
}
