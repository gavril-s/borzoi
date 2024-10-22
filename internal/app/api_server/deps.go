package apiserver

import (
	"context"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
)

type Storage interface {
	UpsertDeploy(ctx context.Context, deploy domain.Deploy) error
	GetDeployByName(ctx context.Context, name string) (*domain.Deploy, error)
	DeleteDeployByName(ctx context.Context, name string) error
}

type DeployManager interface {
	StartDeploy(deploy domain.Deploy, repoPath string) error
}

type PortManager interface {
	FillDeployPorts(ctx context.Context, deploy *domain.Deploy) error
}

type RepoManager interface {
	CloneRepo(repoURL string, branchName string, preparedBranchName string) (string, error)
	ReadBorzoiConfigFromRepo(repoPath string) (*domain.BorzoiConfig, error)
}
