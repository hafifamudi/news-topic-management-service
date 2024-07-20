package request

type CreateNewsRequest struct {
	Title    string   `json:"title" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	Status   string   `json:"status" validate:"required,oneof=draft published"`
	TopicIDs []string `json:"topic_ids"`
}

type UpdateNewsRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Status   string   `json:"status"`
	TopicIDs []string `json:"topic_ids"`
}
