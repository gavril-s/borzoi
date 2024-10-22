package domain

type ServiceConfig struct {
	Name           string `json:"name,omitempty" yaml:"name,omitempty"`
	UpperName      string `json:"upper_name,omitempty" yaml:"upper_name,omitempty"`
	InstanceCount  int    `json:"instance_count,omitempty" yaml:"instance_count,omitempty"`
	LocalOnly      bool   `json:"local_only,omitempty" yaml:"local_only,omitempty"`
	InternalPort   int    `json:"internal_port,omitempty" yaml:"internal_port,omitempty"`
	ExternalPort   int    `json:"external_port,omitempty" yaml:"external_port,omitempty"`
	RootPath       string `json:"root_path,omitempty" yaml:"root_path,omitempty"`
	DockerfilePath string `json:"dockerfile_path,omitempty" yaml:"dockerfile_path,omitempty"`
}
