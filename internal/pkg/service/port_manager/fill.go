package portmanager

import (
	"context"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
)

func (pm *PortManager) FillDeployPorts(ctx context.Context, deploy *domain.Deploy) error {
	var err error
	for i, service := range deploy.Services {
		for j := range service.Nodes {
			deploy.Services[i].Nodes[j].Port, err = pm.allocatePort(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
