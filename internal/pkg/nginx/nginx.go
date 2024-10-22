package nginx

import (
	"fmt"
	"strings"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"github.com/gavril-s/borzoi/internal/pkg/hosts"
)

func buildUpstreamBody(addresses []string) string {
	var upstreamBody strings.Builder
	for _, address := range addresses {
		upstreamBody.WriteString(fmt.Sprintf("\tserver %s;\n", address))
	}
	return upstreamBody.String()
}

func serverName(deploy domain.Deploy) string {
	return fmt.Sprintf("%s www.%s host.docker.internal", deploy.URL, deploy.URL)
}

func BuildNginxConfig(deploy domain.Deploy) string {
	var nginxConfig strings.Builder

	for _, service := range deploy.Services {
		nginxConfig.WriteString(fmt.Sprintf("upstream %s {\n", service.Config.Name))

		addrs := make([]string, 0, len(service.Nodes))
		for _, node := range service.Nodes {
			addrs = append(addrs, fmt.Sprintf("%s:%d", hosts.LocalhostByIP, node.Port))
		}
		nginxConfig.WriteString(buildUpstreamBody(addrs))

		nginxConfig.WriteString("}\n\n")
	}

	for i, service := range deploy.Services {
		nginxConfig.WriteString("server {\n")
		nginxConfig.WriteString(fmt.Sprintf("\tlisten %d;\n", service.Config.ExternalPort))
		nginxConfig.WriteString(fmt.Sprintf("\tserver_name %s;\n\n", serverName(deploy)))
		if service.Config.LocalOnly {
			nginxConfig.WriteString(fmt.Sprintf("\tallow %s;\n", hosts.LocalhostByIP))
			nginxConfig.WriteString("\tdeny all;\n\n")
		}
		nginxConfig.WriteString("\tlocation / {\n")
		nginxConfig.WriteString(fmt.Sprintf("\t\tproxy_pass http://%s;\n", service.Config.Name))
		nginxConfig.WriteString("\t\tproxy_set_header Host $host;\n")
		nginxConfig.WriteString("\t\tproxy_set_header X-Real-IP $remote_addr;\n")
		nginxConfig.WriteString("\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n")
		nginxConfig.WriteString("\t\tproxy_set_header X-Forwarded-Proto $scheme;\n")
		nginxConfig.WriteString("\t}\n")
		nginxConfig.WriteString("}\n")
		if i != len(deploy.Services)-1 {
			nginxConfig.WriteString("\n")
		}
	}

	return nginxConfig.String()
}
