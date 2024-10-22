package apiserver

import (
	"net/http"

	"github.com/gavril-s/borzoi/internal/config"
)

type server struct {
	cfg           config.Config
	storage       Storage
	deployManager DeployManager
	portManager   PortManager
	repoManager   RepoManager
}

func NewServer(
	cfg config.Config,
	storage Storage,
	deployManager DeployManager,
	portManager PortManager,
	repoManager RepoManager,
) *server {
	return &server{
		cfg:           cfg,
		storage:       storage,
		deployManager: deployManager,
		portManager:   portManager,
		repoManager:   repoManager,
	}
}

func (s *server) CreateDeploy(w http.ResponseWriter, r *http.Request) {
	s.postMiddleware(http.HandlerFunc(s.createDeploy), s.cfg.RequestTimeout).ServeHTTP(w, r)
}

func (s *server) DeleteDeploy(w http.ResponseWriter, r *http.Request) {
	s.postMiddleware(http.HandlerFunc(s.deleteDeploy), s.cfg.RequestTimeout).ServeHTTP(w, r)
}

func (s *server) RestartDeploy(w http.ResponseWriter, r *http.Request) {
	s.postMiddleware(http.HandlerFunc(s.restartDeploy), s.cfg.RequestTimeout).ServeHTTP(w, r)
}

func (s *server) ListDeploys(w http.ResponseWriter, r *http.Request) {
	s.getMiddleware(http.HandlerFunc(s.listDeploys), s.cfg.RequestTimeout).ServeHTTP(w, r)
}
