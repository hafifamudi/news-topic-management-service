package common

import (
	"github.com/google/uuid"
	"time"
)

type SuccessWithMessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TopicResource struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	News      []NewsResource `json:"news"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type NewsResource struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Status    string          `json:"status"`
	Topics    []TopicResource `json:"topics"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
