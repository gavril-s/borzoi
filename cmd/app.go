package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gavril-s/borzoi/internal/config"
)

var shutdownSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
}

type app struct {
	cfg           config.Config
	apiServer     APIServer
	websiteServer WebsiteServer
	watcher       Watcher
}

func newApp(
	cfg config.Config,
	apiServer APIServer,
	websiteServer WebsiteServer,
	watcher Watcher,
) *app {
	return &app{
		cfg:           cfg,
		apiServer:     apiServer,
		websiteServer: websiteServer,
		watcher:       watcher,
	}
}

func (a *app) runServices(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		runServers(ctx, a.cfg, a.apiServer, a.websiteServer)
		wg.Done()
	}()

	go func() {
		runWatcher(ctx, a.cfg, a.watcher)
		wg.Done()
	}()

	wg.Wait()
}

func (a *app) run(ctx context.Context, cancel context.CancelFunc) {
	stop := make(chan os.Signal)
	signal.Notify(stop, shutdownSignals...)

	go func() {
		a.runServices(ctx)
	}()

	sig := <-stop
	log.Printf("Received signal: %s", sig)
	log.Println("Shutting down...")
	cancel()
}
