package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.WriteResponse, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEnconder(w).Encode(v)
}
