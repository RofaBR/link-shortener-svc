package service

import (
	"net"
	"net/http"

	"github.com/RofaBR/link-shortener-svc/internal/config"
	"github.com/RofaBR/link-shortener-svc/internal/data"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	storage  data.LinkStorage
	config   config.Config
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router(s.storage)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	db := cfg.DB()
	storage := data.NewPostgresLinkStorage(db.RawDB())

	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		storage:  storage,
		config:   cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
