package deploymanager

import (
	"os"
	"os/exec"
	"path/filepath"

	dockercompose "github.com/gavril-s/borzoi/internal/pkg/docker_compose"
	"github.com/gavril-s/borzoi/internal/pkg/domain"
)

func (d *DeployManager) getDockerComposePathAndRelativeRepoPath(repoPath string) (string, string) {
	return filepath.Split(repoPath)
}

func (d *DeployManager) createDockerCompose(deploy domain.Deploy, repoPath string) error {
	dockerComposePath, relativeRepoPatb := d.getDockerComposePathAndRelativeRepoPath(repoPath)
	dockerCompose := dockercompose.BuildDockerCompose(deploy, relativeRepoPatb)
	return os.WriteFile(filepath.Join(dockerComposePath, "docker-compose.yaml"), []byte(dockerCompose), 0644)
}

func (d *DeployManager) runDockerCompose(borzoiDirPath string) error {
	cmd := exec.Command("docker-compose", "up", "-d", "--build")
	cmd.Dir = borzoiDirPath
	return cmd.Run()
}
