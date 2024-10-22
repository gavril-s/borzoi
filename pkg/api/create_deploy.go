package api

type CreateDeployRequest struct {
	RepoURL    string `json:"repo_url,omitempty"`
	BranchName string `json:"branch_name,omitempty"`
}

type CreateDeployResponse struct {
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}
