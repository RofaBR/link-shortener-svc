package service

import (
	"github.com/RofaBR/link-shortener-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/link-shortener-svc", func(r chi.Router) {
		r.Post("/links", handlers.CreateLink)
		r.Get("/links/{shortURL}", handlers.GetLink)
	})

	return r
}
