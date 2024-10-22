package apiserver

import (
	"encoding/json"
	"net/http"

	branchname "github.com/gavril-s/borzoi/internal/pkg/branch_name"
	"github.com/gavril-s/borzoi/internal/pkg/domain"
	errorwriter "github.com/gavril-s/borzoi/internal/pkg/error_writer"
	"github.com/gavril-s/borzoi/pkg/api"
)

func (s *server) createDeploy(w http.ResponseWriter, r *http.Request) {
	var req api.CreateDeployRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	deployName := branchname.PrepareBranchName(req.BranchName)

	repoPath, err := s.repoManager.CloneRepo(req.RepoURL, req.BranchName, deployName)
	if err != nil {
		errorwriter.WriteError(w, http.StatusInternalServerError, "Failed to clone repo", err)
		return
	}

	config, err := s.repoManager.ReadBorzoiConfigFromRepo(repoPath)
	if err != nil {
		errorwriter.WriteError(w, http.StatusInternalServerError, "Failed to read config from repo", err)
		return
	}

	deploy := domain.NewDeploy(deployName, config, req.RepoURL, req.BranchName)

	err = s.portManager.FillDeployPorts(r.Context(), &deploy)
	if err != nil {
		errorwriter.WriteError(w, http.StatusInternalServerError, "Failed to allocate ports for deploy", err)
		return
	}

	err = s.storage.UpsertDeploy(r.Context(), deploy)
	if err != nil {
		errorwriter.WriteError(w, http.StatusInternalServerError, "Failed to upsert depoloy", err)
		return
	}

	go s.deployManager.StartDeploy(deploy, repoPath)

	resp := api.CreateDeployResponse{
		Name:   deploy.Name,
		Status: string(deploy.Status),
	}
	json.NewEncoder(w).Encode(resp)
}
