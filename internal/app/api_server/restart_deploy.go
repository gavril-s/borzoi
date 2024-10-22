package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gavril-s/borzoi/pkg/api"
)

func (s *server) restartDeploy(w http.ResponseWriter, r *http.Request) {
	var req api.RestartDeployRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	resp := api.RestartDeployResponse{}
	json.NewEncoder(w).Encode(resp)
}
