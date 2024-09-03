package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	response := Response{
		Status:  "error",
		Message: message,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ParseID(r *http.Request) (int64, error) {
	idStr := r.URL.Query().Get("id")
	return strconv.ParseInt(idStr, 10, 64)
}
