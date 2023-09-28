package router

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"go-cloud-native-rest-api/api/resource/book"
	"go-cloud-native-rest-api/api/resource/health"
)

func New(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookAPI := book.New(db)
		r.Get("/books", bookAPI.List)
		r.Post("/books", bookAPI.Create)
		r.Get("/books/{id}", bookAPI.Read)
		r.Put("/books/{id}", bookAPI.Update)
		r.Delete("/books/{id}", bookAPI.Delete)
	})

	return r
}
