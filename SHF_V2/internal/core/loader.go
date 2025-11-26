package core

import (
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

type Registry struct {
    ByID    map[string]Module
    ByAlias map[string]Module
    All     []Module
}

func LoadRegistry() Registry {
    exe, err := os.Executable()
    if err != nil {
        fmt.Println("[!] Could not determine executable path:", err)
        return Registry{ByID: map[string]Module{}, ByAlias: map[string]Module{}, All: []Module{}}
    }
    root := filepath.Dir(exe)
    modulesRoot := filepath.Join(root, "modules")

    reg := Registry{
        ByID:    map[string]Module{},
        ByAlias: map[string]Module{},
        All:     []Module{},
    }

    filepath.Walk(modulesRoot, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return nil
        }
        if info.IsDir() {
            return nil
        }
        if info.Name() != "module.yaml" {
            return nil
        }

        data, err := os.ReadFile(path)
        if err != nil {
            fmt.Println("[!] Failed to read module.yaml:", err)
            return nil
        }

        var m Module
        if err := yaml.Unmarshal(data, &m); err != nil {
            fmt.Println("[!] Failed to parse module.yaml:", err)
            return nil
        }
        m.BaseDir = filepath.Dir(path)
        reg.ByID[m.ID] = m
        reg.All = append(reg.All, m)
        for _, a := range m.Aliases {
            reg.ByAlias[a] = m
        }

        return nil
    })

    return reg
}
