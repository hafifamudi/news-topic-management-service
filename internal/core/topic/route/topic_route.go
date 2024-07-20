package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/controller"
)

func Register(r chi.Router) {
	topicController := controller.Topic()

	r.Route("/topics", func(r chi.Router) {
		r.Get("/", topicController.ListTopic)
		r.Get("/{id}", topicController.DetailTopic)
		r.Post("/", topicController.CreateTopic)
		r.Put("/{id}", topicController.UpdateTopic)
		r.Delete("/{id}", topicController.DeleteTopic)
	})

}
