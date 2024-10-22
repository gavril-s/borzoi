package dockercompose

import (
	"fmt"
	"path"
	"strings"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
)

func indent(level int) string {
	var res strings.Builder
	for range level {
		res.WriteString("  ") // 2 spaces
	}
	return res.String()
}

func BuildDockerCompose(deploy domain.Deploy, relativeRepoPath string) string {
	var dockerCompose strings.Builder

	dockerCompose.WriteString("services:\n")
	for _, service := range deploy.Services {
		for i, node := range service.Nodes {
			dockerCompose.WriteString(fmt.Sprintf("%s%s-%d:\n", indent(1), service.Config.Name, i))
			dockerCompose.WriteString(fmt.Sprintf("%sbuild:\n", indent(2)))
			dockerCompose.WriteString(fmt.Sprintf("%scontext: %s\n", indent(3), path.Join(".", relativeRepoPath, service.Config.RootPath)))
			dockerCompose.WriteString(fmt.Sprintf("%sdockerfile: %s\n", indent(3), service.Config.DockerfilePath))
			dockerCompose.WriteString(fmt.Sprintf("%senvironment:\n", indent(2)))
			for _, service := range deploy.Services {
				dockerCompose.WriteString(fmt.Sprintf("%s- %s_URL=http://host.docker.internal:%d\n", indent(3), service.Config.UpperName, service.Config.ExternalPort))
			}
			dockerCompose.WriteString(fmt.Sprintf("%sports:\n", indent(2)))
			dockerCompose.WriteString(fmt.Sprintf("%s- \"%d:%d\"\n", indent(3), node.Port, service.Config.InternalPort))
		}
	}

	return dockerCompose.String()
}
