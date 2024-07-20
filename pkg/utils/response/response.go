package response

import (
	"encoding/json"
	"net/http"
)

type SuccessWithMessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorValidationResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func SuccessWithMessage(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessWithMessageResponse{
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
	})
}

func ErrorValidation(w http.ResponseWriter, errors interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(ErrorValidationResponse{
		Message: "The given data was invalid.",
		Errors:  errors,
	})
}
