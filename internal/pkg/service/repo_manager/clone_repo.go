package repomanager

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (m *RepoManager) getRepoName(repoURL string) string {
	parts := strings.Split(repoURL, "/")
	if len(parts) == 0 {
		return repoURL
	}
	return parts[len(parts)-1]
}

func (m *RepoManager) cleanDirPath(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error cleaning up for cloning: %v", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("error cleaning up for cloning: %s is not a directory", dirPath)
	}
	return os.RemoveAll(dirPath)
}

func (m *RepoManager) CloneRepo(repoURL string, branchName string, preparedBranchName string) (string, error) {
	repoPath := m.getRepoPathByURLAndPreparedBranchName(repoURL, preparedBranchName)
	if err := m.cleanDirPath(repoPath); err != nil {
		return "", err
	}
	cmd := exec.Command("git", "clone", "-b", branchName, repoURL, repoPath)
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return repoPath, nil
}
