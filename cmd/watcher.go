package main

import (
	"context"

	"github.com/gavril-s/borzoi/internal/config"
)

type Watcher interface {
	WatchDeploys()
}

func runWatcher(ctx context.Context, cfg config.Config, watcher Watcher) {
	if cfg.EnableWatcher {
		watcher.WatchDeploys()
	}
}
