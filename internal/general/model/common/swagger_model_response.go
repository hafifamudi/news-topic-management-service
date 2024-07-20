package common

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type SuccessWithMessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessWithMessage(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessWithMessageResponse{
		Message: message,
		Data:    data,
	})
}

// TopicResource represents a topic with its associated news.
type TopicResource struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	News      []NewsResource `json:"news"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// NewsResource represents a news item with its associated topics.
type NewsResource struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Status    string          `json:"status"`
	Topics    []TopicResource `json:"topics"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
