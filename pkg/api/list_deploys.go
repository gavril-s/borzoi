package api

type Deploy struct {
}

type ListDeploysResponse struct {
	Deploys []Deploy `json:"deploys"`
}
