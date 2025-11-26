package main

import (
    "fmt"
    "os"

    "shf/internal/cli"
    "shf/internal/core"
)

func main() {
    cfg := core.LoadConfig()
    logger := core.NewLogger(cfg)

    if len(os.Args) == 1 {
        cli.PrintBanner()
        return
    }

    registry := core.LoadRegistry()

    cmd := os.Args[1]
    args := os.Args[2:]

    switch cmd {
    case "help", "--help", "-h":
        cli.PrintHelp()
        return
    case "list":
        cli.ListModules(registry)
        return
    case "search":
        cli.SearchModules(registry, args)
        return
    case "info":
        cli.InfoModule(registry, args)
        return
    case "run":
        cli.RunModuleExplicit(registry, cfg, logger, args)
        return
    case "update":
        cli.UpdateSHF(cfg)
        return
    case "new", "scaffold":
        cli.ScaffoldModule()
        return
    case "version":
        fmt.Println("SHF - SudoSoc Hybrid Framework version 0.5.0")
        return
    }

    // Otherwise: treat as alias / smart execution
    cli.RunModuleSmart(registry, cfg, logger, cmd, args)
}
