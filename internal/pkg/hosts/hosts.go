package hosts

import (
	"fmt"
	"os"
	"strings"

	"github.com/gavril-s/borzoi/internal/config"
)

type Host string

const (
	LocalhostByIP   Host = "127.0.0.1"
	LocalhostByName Host = "localhost"
)

func AppendDomainToHosts(cfg config.Config, domain string) error {
	data, err := os.ReadFile(cfg.HostsFilePath)
	if err != nil {
		return fmt.Errorf("error reading hosts file: %w", err)
	}

	if strings.Contains(string(data), domain) {
		return nil
	}

	newData := fmt.Sprintf("\n%s %s www.%s\n", LocalhostByIP, domain, domain)
	err = os.WriteFile(cfg.HostsFilePath, append(data, newData...), 0644)
	if err != nil {
		return fmt.Errorf("error writing to hosts file: %w", err)
	}
	return nil
}
