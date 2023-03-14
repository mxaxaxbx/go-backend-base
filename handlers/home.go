package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mxaxaxbx/go-backend-base/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.MainResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    make(map[string]string),
	})
	fmt.Println("OK")
}
