package deploymanager

import (
	"os"
	"os/exec"
	"path"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"github.com/gavril-s/borzoi/internal/pkg/nginx"
)

func (d *DeployManager) createNginxConfig(deploy domain.Deploy) error {
	nginxConfig := nginx.BuildNginxConfig(deploy)
	nginxConfigPath := path.Join(d.cfg.NginxSitesAvailablePath, deploy.Name)
	return os.WriteFile(nginxConfigPath, []byte(nginxConfig), 0644)
}

func (d *DeployManager) applyNginxConfig(deploy domain.Deploy) error {
	nginxConfigPath := path.Join(d.cfg.NginxSitesAvailablePath, deploy.Name)
	symlinkPath := path.Join(d.cfg.NginxSitesEnabledPath, deploy.Name)

	if _, err := os.Stat(symlinkPath); os.IsNotExist(err) {
		err = os.Symlink(nginxConfigPath, symlinkPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DeployManager) restartNginx() error {
	cmd := exec.Command("systemctl", "restart", "nginx")
	return cmd.Run()
}
