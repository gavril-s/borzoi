package main

import (
	"context"

	apiserver "github.com/gavril-s/borzoi/internal/app/api_server"
	"github.com/gavril-s/borzoi/internal/app/watcher"
	websiteserver "github.com/gavril-s/borzoi/internal/app/website_server"
	"github.com/gavril-s/borzoi/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := config.NewConfig()
	deps := buildDeps(ctx, cfg)

	api := apiserver.NewServer(cfg, deps.db, deps.deployManager, deps.portManager, deps.repoManager)
	website := websiteserver.NewServer(cfg, deps.db, deps.pageBuilder)
	watcher := watcher.NewWatcher(cfg, deps.db)

	app := newApp(cfg, api, website, watcher)
	app.run(ctx, cancel)
}
