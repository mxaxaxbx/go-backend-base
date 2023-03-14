package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mxaxaxbx/go-backend-base/models"
)

func MeHandler(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"id":    0,
		"name":  "Miguel A. Arenas",
		"email": "mxaxaxbx@protonmail.com",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.MainResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    m,
	})
}
