package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"rest-with-dynamodb/handler/book"
	"rest-with-dynamodb/handler/health"
)

func (s *Server) httpRouter() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	//middleware
	//recover panics
	s.router.Use(middleware.Recoverer)

	s.router.Get("/health", health.Handler)
	s.router.Route("/api/book", func(r chi.Router) {
		r.Put("/", book.AddBookHandler)
		r.Route("/{isbn}", func(r chi.Router) {
			r.Get("/", book.GetBookHandler)
		})
	})
}
