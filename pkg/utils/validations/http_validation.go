package validations

import (
	"encoding/json"
	"github.com/hafifamudi/news-topic-management-service/pkg/utils/errors"
	"net/http"
)

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func HandleHttpRequestValidationError(w http.ResponseWriter, err error) {
	formattedErrors := errors.FormatValidationError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Errors: formattedErrors})
}
