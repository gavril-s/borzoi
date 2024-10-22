package repomanager

import (
	"path/filepath"

	"github.com/gavril-s/borzoi/internal/config"
)

type RepoManager struct {
	cfg config.Config
}

func NewRepoManager(cfg config.Config) *RepoManager {
	return &RepoManager{cfg: cfg}
}

func (m *RepoManager) getRepoPathByURLAndPreparedBranchName(repoURL string, preparedBranchName string) string {
	return filepath.Join(m.cfg.RuntimeDirPath, "repos", m.getRepoName(repoURL), preparedBranchName, "repo")
}
