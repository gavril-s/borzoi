package main

import (
	"context"

	"github.com/gavril-s/borzoi/internal/config"
	deploymanager "github.com/gavril-s/borzoi/internal/pkg/service/deploy_manager"
	pagebuilder "github.com/gavril-s/borzoi/internal/pkg/service/page_builder"
	portmanager "github.com/gavril-s/borzoi/internal/pkg/service/port_manager"
	repomanager "github.com/gavril-s/borzoi/internal/pkg/service/repo_manager"
	"github.com/gavril-s/borzoi/internal/pkg/storage"
)

type deps struct {
	db            *storage.Redis
	deployManager *deploymanager.DeployManager
	pageBuilder   *pagebuilder.PageBuilder
	portManager   *portmanager.PortManager
	repoManager   *repomanager.RepoManager
}

func buildDeps(ctx context.Context, cfg config.Config) *deps {
	db := config.ConnectToRedis(ctx, cfg)
	return &deps{
		db:            db,
		deployManager: deploymanager.NewDeployManager(cfg),
		pageBuilder:   pagebuilder.NewPageBuilder(),
		portManager:   portmanager.NewPortManager(cfg, db, [2]int{cfg.DefaultPortRangeMin, cfg.DefaultPortRangeMax}),
		repoManager:   repomanager.NewRepoManager(cfg),
	}
}
