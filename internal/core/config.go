package core

import (
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

type Config struct {
    RepoURL    string `yaml:"repo_url"`
    RepoBranch string `yaml:"repo_branch"`
    Colors     bool   `yaml:"colors"`
    LogEnabled bool   `yaml:"log_enabled"`
    JSONIndent int    `yaml:"json_indent"`
    Cache      bool   `yaml:"cache"`
    AutoUpdate bool   `yaml:"auto_update"`
}

func defaultConfig() Config {
    return Config{
        RepoURL:    "https://github.com/sudosoc/SHF.git",
        RepoBranch: "main",
        Colors:     true,
        LogEnabled: true,
        JSONIndent: 2,
        Cache:      true,
        AutoUpdate: false,
    }
}

func configPath() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    dir := filepath.Join(home, ".shf")
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        if mkErr := os.MkdirAll(dir, 0o755); mkErr != nil {
            return "", mkErr
        }
    }
    return filepath.Join(dir, "config.yaml"), nil
}

func LoadConfig() Config {
    cfg := defaultConfig()

    path, err := configPath()
    if err != nil {
        fmt.Println("[!] Could not determine config path:", err)
        return cfg
    }

    data, err := os.ReadFile(path)
    if err != nil {
        if writeErr := writeDefaultConfig(path, cfg); writeErr != nil {
            fmt.Println("[!] Could not write default config:", writeErr)
        }
        return cfg
    }

    if err := yaml.Unmarshal(data, &cfg); err != nil {
        fmt.Println("[!] Failed to parse config, using defaults:", err)
        return defaultConfig()
    }

    return cfg
}

func writeDefaultConfig(path string, cfg Config) error {
    b, err := yaml.Marshal(&cfg)
    if err != nil {
        return err
    }
    return os.WriteFile(path, b, 0o644)
}
