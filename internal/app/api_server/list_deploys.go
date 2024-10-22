package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gavril-s/borzoi/pkg/api"
)

func (s *server) listDeploys(w http.ResponseWriter, r *http.Request) {
	resp := api.ListDeploysResponse{}
	json.NewEncoder(w).Encode(resp)
}
