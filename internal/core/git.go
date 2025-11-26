package core

import (
    "os"
    "os/exec"
    "path/filepath"
)

func NewGitPullCommand(root string, branch string) *exec.Cmd {
    if _, err := exec.LookPath("git"); err != nil {
        return nil
    }

    gitDir := filepath.Join(root, ".git")
    if _, err := os.Stat(gitDir); err != nil {
        return nil
    }

    cmd := exec.Command("git", "-C", root, "pull", "origin", branch)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd
}
