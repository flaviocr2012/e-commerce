package utils

import (
	"encoding/json"
	"net/http"
	"strconv" // Adicione esta linha
	"github.com/gorilla/mux"
)

// ParseID extrai o ID da URL path
func ParseID(r *http.Request) (int64, error) {
	// Usa o gorilla/mux para extrair variáveis da URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, &ParseError{Message: "ID not found in URL"}
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, &ParseError{Message: "Invalid ID format"}
	}

	return id, nil
}

// HandleError envia uma resposta de erro JSON
func HandleError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: message,
	})
}

// Response estrutura padrão para respostas da API
type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// ParseError erro de parsing
type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}