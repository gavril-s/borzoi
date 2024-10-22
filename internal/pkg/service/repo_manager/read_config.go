package repomanager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"gopkg.in/yaml.v3"
)

type borzoiConfig struct {
	ProdBranch string                          `json:"prod_branch,omitempty" yaml:"prod_branch,omitempty"`
	BaseURL    string                          `json:"base_url,omitempty" yaml:"base_url,omitempty"`
	Services   map[string]domain.ServiceConfig `json:"services,omitempty" yaml:"services,omitempty"`
}

func (m *RepoManager) ReadBorzoiConfigFromRepo(repoPath string) (*domain.BorzoiConfig, error) {
	path := filepath.Join(repoPath, m.cfg.BorzoiConfigPath)

	yamlData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var config borzoiConfig
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal borzoi config file: %v", err)
	}

	res := domain.BorzoiConfig{
		ProdBranch: config.ProdBranch,
		BaseURL:    config.BaseURL,
		Services:   make([]domain.ServiceConfig, 0, len(config.Services)),
	}
	for name, service := range config.Services {
		service.Name = name
		service.UpperName = strings.ToUpper(name)
		res.Services = append(res.Services, service)
	}

	return &res, nil
}
