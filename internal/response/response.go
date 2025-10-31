package response

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	response := JsonResponse{
		Message: "success",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	response := JsonResponse{
		Message: message,
		Error:   err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func CreateResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, err error) {
	response := JsonResponse{
		Message: message,
		Data:    data,
		Error:   err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
