package routes

import (
	"github.com/go-chi/chi/v5"
	"news-topic-management-service/internal/core/news/controller"
)

func Register(r chi.Router) {
	newsController := controller.News()

	r.Route("/news", func(r chi.Router) {
		r.Get("/", newsController.ListNews)
		r.Get("/{id}", newsController.DetailNews)
		r.Post("/", newsController.CreateNews)
		r.Put("/{id}", newsController.UpdateNews)
		r.Delete("/{id}", newsController.DeleteNews)
	})
}
