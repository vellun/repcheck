package cloner

import (
	"fmt"
	"os"
	"os/exec"
)

type RepoCloner interface {
	Clone(repoURL string) (string, error)
}

type GitCloner struct{}

func (g *GitCloner) Clone(repoURL string, pwd string) (string, error) {
	repoDir, err := os.MkdirTemp(pwd, "go-repo-")

	if err != nil {
		return "", fmt.Errorf("Repo cloning error: %v", err)
	}

	fmt.Println("Getting repo...")
	cloneCmd := exec.Command("git", "clone", repoURL, repoDir)
	if err := cloneCmd.Run(); err != nil {
		return "", fmt.Errorf("Repo cloning error: %v", err)
	}
	return repoDir, nil
}
