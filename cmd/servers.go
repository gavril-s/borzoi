package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gavril-s/borzoi/internal/config"
)

type WebsiteServer interface {
	ServeIndex(w http.ResponseWriter, r *http.Request)
}

type APIServer interface {
	CreateDeploy(w http.ResponseWriter, r *http.Request)
	DeleteDeploy(w http.ResponseWriter, r *http.Request)
	RestartDeploy(w http.ResponseWriter, r *http.Request)
	ListDeploys(w http.ResponseWriter, r *http.Request)
}

func handlerFuncWithCtx(ctx context.Context, serve func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serve(w, r.WithContext(ctx))
	})
}

func runServers(ctx context.Context, cfg config.Config, apiServer APIServer, websiteServer WebsiteServer) {
	http.Handle("/", handlerFuncWithCtx(ctx, websiteServer.ServeIndex))

	http.Handle("/api/v1/deploy/create", handlerFuncWithCtx(ctx, apiServer.CreateDeploy))
	http.Handle("/api/v1/deploy/delete", handlerFuncWithCtx(ctx, apiServer.DeleteDeploy))
	http.Handle("/api/v1/deploy/restart", handlerFuncWithCtx(ctx, apiServer.RestartDeploy))
	http.Handle("/api/v1/deploys", handlerFuncWithCtx(ctx, apiServer.ListDeploys))

	serverAddress := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Server started at %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
