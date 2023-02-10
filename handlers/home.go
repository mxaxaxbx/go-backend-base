package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mxaxaxbx/go-backend-base/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.MainResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    make(map[string]string),
	})
	fmt.Println("OK")
}
