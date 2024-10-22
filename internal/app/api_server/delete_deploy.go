package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gavril-s/borzoi/pkg/api"
)

func (s *server) deleteDeploy(w http.ResponseWriter, r *http.Request) {
	var req api.DeleteDeployRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	resp := api.DeleteDeployResponse{}
	json.NewEncoder(w).Encode(resp)
}
