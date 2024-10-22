package watcher

import (
	"github.com/gavril-s/borzoi/internal/config"
)

type watcher struct {
}

func NewWatcher(cfg config.Config, storage Storage) *watcher {
	return &watcher{}
}
