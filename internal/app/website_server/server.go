package websiteserver

import (
	"net/http"

	"github.com/gavril-s/borzoi/internal/config"
)

type server struct {
	cfg         config.Config
	storage     Storage
	pageBuilder PageBuilder
}

func NewServer(cfg config.Config, storage Storage, pageBuilder PageBuilder) *server {
	return &server{
		cfg:         cfg,
		storage:     storage,
		pageBuilder: pageBuilder,
	}
}

func (s *server) ServeIndex(w http.ResponseWriter, r *http.Request) {
	s.middleware(http.HandlerFunc(s.serveIndex), s.cfg.RequestTimeout).ServeHTTP(w, r)
}
