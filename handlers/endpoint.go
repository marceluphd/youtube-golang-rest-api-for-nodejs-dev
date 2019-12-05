package handlers

import (
	"github.com/go-chi/chi"
)

func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", s.registerUser())
		})

		r.Route("/todos", func(r chi.Router) {
			r.Use(s.withUser)
			r.Post("/", s.createTodo())

			r.Route("/{id}", func(r chi.Router) {
				r.Use(s.todoCtx)
				r.Use(s.withOwner("todo"))
				r.Patch("/", s.updateTodo())
				r.Delete("/", s.deleteTodo())
			})
		})
	})
}
