package deploymanager

import "github.com/gavril-s/borzoi/internal/config"

type DeployManager struct {
	cfg config.Config
}

func NewDeployManager(cfg config.Config) *DeployManager {
	return &DeployManager{cfg: cfg}
}
