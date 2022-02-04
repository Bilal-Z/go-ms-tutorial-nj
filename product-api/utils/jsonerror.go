package utils

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, err CustomError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}