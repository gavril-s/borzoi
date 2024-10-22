package domain

type BorzoiConfig struct {
	ProdBranch string          `json:"prod_branch,omitempty" yaml:"prod_branch,omitempty"`
	BaseURL    string          `json:"base_url,omitempty" yaml:"base_url,omitempty"`
	Services   []ServiceConfig `json:"services,omitempty" yaml:"services,omitempty"`
}
