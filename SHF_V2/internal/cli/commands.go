package cli

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "shf/internal/core"
)

func ListModules(reg core.Registry) {
    fmt.Println("================= MODULES =================")
    groups := map[string][]core.Module{}
    for _, m := range reg.All {
        cat := m.ID
        if idx := strings.Index(cat, "/"); idx != -1 {
            cat = cat[:idx]
        }
        groups[cat] = append(groups[cat], m)
    }

    order := []string{"offensive", "defensive", "forensics", "threat_intel"}
    for _, cat := range order {
        mods, ok := groups[cat]
        if !ok {
            continue
        }
        fmt.Println()
        fmt.Println("[ " + strings.Title(strings.ReplaceAll(cat, "_", " ")) + " ]")
        for _, m := range mods {
            fmt.Println("  " + m.ID)
            fmt.Println("      - " + m.Name)
        }
    }
}

func SearchModules(reg core.Registry, args []string) {
    if len(args) == 0 {
        fmt.Println("[!] Please provide a search term.")
        return
    }
    term := strings.ToLower(strings.Join(args, " "))

    fmt.Println("Search results for:", term)
    found := false
    for _, m := range reg.All {
        text := strings.ToLower(m.ID + " " + m.Name + " " + strings.Join(m.Tags, " ") + " " + strings.Join(m.Aliases, " "))
        if strings.Contains(text, term) {
            fmt.Println("  " + m.ID)
            fmt.Println("      - " + m.Name)
            fmt.Println("      aliases:", strings.Join(m.Aliases, ", "))
            found = true
        }
    }
    if !found {
        fmt.Println("  No modules matched.")
    }
}

func InfoModule(reg core.Registry, args []string) {
    if len(args) == 0 {
        fmt.Println("[!] Please provide a module id or alias.")
        return
    }
    key := args[0]

    m, ok := reg.ByID[key]
    if !ok {
        if m2, ok2 := reg.ByAlias[key]; ok2 {
            m = m2
        } else {
            fmt.Println("[!] Module not found:", key)
            return
        }
    }

    fmt.Println("ID:      ", m.ID)
    fmt.Println("Name:    ", m.Name)
    fmt.Println("Type:    ", m.Type)
    fmt.Println("Path:    ", m.Path)
    fmt.Println("Aliases: ", strings.Join(m.Aliases, ", "))
    fmt.Println("Tags:    ", strings.Join(m.Tags, ", "))
    fmt.Println()

    if m.HelpFile != "" {
        helpPath := filepath.Join(m.BaseDir, m.HelpFile)
        if data, err := os.ReadFile(helpPath); err == nil {
            fmt.Println(string(data))
        } else {
            fmt.Println("[!] Could not read help file:", err)
        }
    }
}

func RunModuleExplicit(reg core.Registry, cfg core.Config, logger core.Logger, args []string) {
    if len(args) == 0 {
        fmt.Println("[!] Please provide a module id.")
        return
    }
    id := args[0]
    m, ok := reg.ByID[id]
    if !ok {
        fmt.Println("[!] Module not found:", id)
        return
    }
    core.RunModule(m, cfg, logger, args[1:], false)
}

func RunModuleSmart(reg core.Registry, cfg core.Config, logger core.Logger, alias string, args []string) {
    if m, ok := reg.ByAlias[alias]; ok {
        core.RunModule(m, cfg, logger, args, true)
        return
    }

    if m, ok := reg.ByID[alias]; ok {
        core.RunModule(m, cfg, logger, args, true)
        return
    }

    fmt.Println("[!] Unknown command or module alias:", alias)
    PrintHelp()
}

func UpdateSHF(cfg core.Config) {
    exe, err := os.Executable()
    if err != nil {
        fmt.Println("[!] Could not determine executable path:", err)
        return
    }
    root := filepath.Dir(exe)

    cmd := core.NewGitPullCommand(root, cfg.RepoBranch)
    if cmd == nil {
        fmt.Println("[!] Git not available or not a git repository.")
        return
    }
    if err := cmd.Run(); err != nil {
        fmt.Println("[!] git pull failed:", err)
    } else {
        fmt.Println("[+] SHF updated from Git.")
    }
}

func ScaffoldModule() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("SHF module scaffolder")
    fmt.Print("Enter module id (e.g. forensics/new_tool): ")
    idLine, _ := reader.ReadString('\n')
    idLine = strings.TrimSpace(idLine)
    if idLine == "" {
        fmt.Println("[!] No id provided.")
        return
    }

    fmt.Print("Language (python/bash): ")
    langLine, _ := reader.ReadString('\n')
    langLine = strings.TrimSpace(langLine)
    if langLine == "" {
        langLine = "python"
    }

    exe, err := os.Executable()
    if err != nil {
        fmt.Println("[!] Could not determine executable path:", err)
        return
    }
    root := filepath.Dir(exe)
    modulesRoot := filepath.Join(root, "modules")

    parts := strings.Split(idLine, "/")
    subPath := filepath.Join(parts...)
    moduleDir := filepath.Join(modulesRoot, subPath)

    if err := os.MkdirAll(moduleDir, 0o755); err != nil {
        fmt.Println("[!] Could not create module directory:", err)
        return
    }

    moduleYamlPath := filepath.Join(moduleDir, "module.yaml")
    if _, err := os.Stat(moduleYamlPath); err == nil {
        fmt.Println("[!] module.yaml already exists, aborting.")
        return
    }

    name := parts[len(parts)-1]

    yamlContent := "id: \"" + idLine + "\"\n" +
        "name: \"" + strings.Title(strings.ReplaceAll(name, "_", " ")) + "\"\n" +
        "type: \"" + langLine + "\"\n" +
        "path: \"" + name + ".py\"\n" +
        "\naliases:\n" +
        "  - \"" + name + "\"\n" +
        "\ntags:\n" +
        "  - \"" + parts[0] + "\"\n" +
        "\nprimary_arg: \"\"\n" +
        "help_file: \"help.txt\"\n"

    if err := os.WriteFile(moduleYamlPath, []byte(yamlContent), 0o644); err != nil {
        fmt.Println("[!] Could not write module.yaml:", err)
        return
    }

    helpPath := filepath.Join(moduleDir, "help.txt")
    helpContent := "Module: " + idLine + "\n" +
        "Description: " + name + " (TODO)\n"
    os.WriteFile(helpPath, []byte(helpContent), 0o644)

    if langLine == "python" {
        scriptPath := filepath.Join(moduleDir, name+".py")
        script := "#!/usr/bin/env python3\n" +
            "import argparse\n\n" +
            "def main():\n" +
            "    parser = argparse.ArgumentParser(description=\"" + name + " module.\")\n" +
            "    # TODO: add arguments here\n" +
            "    args = parser.parse_args()\n" +
            "    print(\"Module " + idLine + " executed.\")\n\n" +
            "if __name__ == \"__main__\":\n" +
            "    main()\n"
        os.WriteFile(scriptPath, []byte(script), 0o755)
    }

    fmt.Println("[+] Module scaffold created at:", moduleDir)
}
