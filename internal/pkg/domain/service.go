package domain

import "fmt"

type Service struct {
	Config ServiceConfig `json:"service_config,omitempty"`
	Nodes  []Node        `json:"nodes,omitempty"`
}

func (s *Service) URL(deploy Deploy) string {
	return fmt.Sprintf("%s:%d", deploy.URL, s.Config.ExternalPort)
}
