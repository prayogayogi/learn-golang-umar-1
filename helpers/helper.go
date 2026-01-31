package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func JsonResponse(w http.ResponseWriter, status int, message string, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := Response{
		Status: status,
		Message: message,
		Data: data,
	}
	json.NewEncoder(w).Encode(response)
}
