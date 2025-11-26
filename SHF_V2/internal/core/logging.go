package core

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"
)

type Logger struct {
    enabled bool
    path    string
}

func NewLogger(cfg Config) Logger {
    if !cfg.LogEnabled {
        return Logger{enabled: false}
    }
    home, err := os.UserHomeDir()
    if err != nil {
        return Logger{enabled: false}
    }
    dir := filepath.Join(home, ".shf", "logs")
    os.MkdirAll(dir, 0o755)
    path := filepath.Join(dir, "shf.log")
    return Logger{
        enabled: true,
        path:    path,
    }
}

func (l Logger) Log(command string, args []string, result error) {
    if !l.enabled {
        return
    }
    f, err := os.OpenFile(l.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
    if err != nil {
        return
    }
    defer f.Close()

    ts := time.Now().Format(time.RFC3339)
    line := fmt.Sprintf("%s | %s %s", ts, command, strings.Join(args, " "))
    if result != nil {
        line += fmt.Sprintf(" | ERROR: %v", result)
    }
    line += "\n"
    f.WriteString(line)
}
