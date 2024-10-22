package domain

import (
	"fmt"
)

type DeployStatus string

const (
	DeployStatusCreated DeployStatus = "created"
	DeployStatusRunning DeployStatus = "running"
	DeployStatusStopped DeployStatus = "stopped"
	DeployStatusFailed  DeployStatus = "failed"
)

type Deploy struct {
	Name     string       `json:"name,omitempty"`
	Services []Service    `json:"services,omitempty"`
	URL      string       `json:"url,omitempty"`
	ReploURL string       `json:"repo_url,omitempty"`
	Status   DeployStatus `json:"status,omitempty"`
}

func NewDeploy(name string, config *BorzoiConfig, repoURL string, branchName string) Deploy {
	deploy := Deploy{
		Name:     name,
		ReploURL: repoURL,
		Status:   DeployStatusCreated,
	}

	if branchName == config.ProdBranch {
		deploy.URL = fmt.Sprintf("%s.%s", name, config.BaseURL)
	} else {
		deploy.URL = config.BaseURL
	}

	deploy.Services = make([]Service, 0, len(config.Services))
	for _, serviceConfig := range config.Services {
		deploy.Services = append(deploy.Services, Service{
			Config: serviceConfig,
			Nodes:  make([]Node, serviceConfig.InstanceCount),
		})
	}

	return deploy
}
