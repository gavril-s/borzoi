package portmanager

import (
	"context"
	"fmt"
	"math/rand"
	"net"
)

func isPortAvailable(port int) bool {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	listener.Close()
	return true
}

func (pm *PortManager) getRandomPort() int {
	return rand.Intn(pm.portRange[1]-pm.portRange[0]) + pm.portRange[0]
}

func (pm *PortManager) allocatePort(ctx context.Context) (int, error) {
	for range pm.cfg.MaxPortSearchIterations {
		port := pm.getRandomPort()

		markedAsBusy, err := pm.storage.IsPortBusy(ctx, port)
		if err != nil {
			return 0, fmt.Errorf("Storage.IsPortBusy: %v", err)
		}

		if !markedAsBusy && isPortAvailable(port) {
			if err := pm.storage.MarkPortAsBusy(ctx, port); err != nil {
				return 0, fmt.Errorf("Storage.MarkPortAsBusy: %v", err)
			}
			return port, nil
		}
	}
	return 0, fmt.Errorf("MaxPortSearchIterations limit exceeded")
}
