package portmanager

import (
	"github.com/gavril-s/borzoi/internal/config"
)

type PortManager struct {
	storage   Storage
	cfg       config.Config
	portRange [2]int
}

func NewPortManager(cfg config.Config, storage Storage, portRange [2]int) *PortManager {
	return &PortManager{
		cfg:       cfg,
		storage:   storage,
		portRange: portRange,
	}
}
