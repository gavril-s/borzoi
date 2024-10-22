package deploymanager

import (
	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"github.com/gavril-s/borzoi/internal/pkg/hosts"
)

func (d *DeployManager) StartDeploy(deploy domain.Deploy, repoPath string) error {
	if err := d.setupDeploy(deploy, repoPath); err != nil {
		return err
	}
	return d.runDeploy(repoPath)
}

func (d *DeployManager) setupDeploy(deploy domain.Deploy, repoPath string) error {
	var err error

	hosts.AppendDomainToHosts(d.cfg, deploy.URL)

	err = d.createDockerCompose(deploy, repoPath)
	if err != nil {
		return err
	}

	err = d.createNginxConfig(deploy)
	if err != nil {
		return err
	}

	err = d.applyNginxConfig(deploy)
	if err != nil {
		return err
	}

	return nil
}

func (d *DeployManager) runDeploy(repoPath string) error {
	var err error

	err = d.runDockerCompose(repoPath)
	if err != nil {
		return err
	}
	err = d.restartNginx()
	if err != nil {
		return err
	}

	return err
}
