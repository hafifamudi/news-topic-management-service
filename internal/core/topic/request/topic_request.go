package request

type CreateTopicRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type UpdateTopicRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
