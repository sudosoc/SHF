package cli

import (
    "fmt"
    "os"
    "path/filepath"
)

func PrintBanner() {
    exe, err := os.Executable()
    if err == nil {
        root := filepath.Dir(exe)
        bannerPath := filepath.Join(root, "docs", "banner.txt")
        if data, err2 := os.ReadFile(bannerPath); err2 == nil {
            fmt.Print(string(data))
            return
        }
    }

    fmt.Println("SHF - SudoSoc Hybrid Framework (v0.2.0)")
    fmt.Println("Use: shf help | shf list | shf search | shf run <id>")
}

func PrintHelp() {
    fmt.Println("SHF - SudoSoc Hybrid Framework (v0.5.0)")
    fmt.Println()
    fmt.Println("Usage:")
    fmt.Println("  shf [command] [options]")
    fmt.Println()
    fmt.Println("Commands:")
    fmt.Println("  list                    List all available modules")
    fmt.Println("  search <term>           Search modules by id/name/tags/aliases")
    fmt.Println("  info <id-or-alias>      Show detailed info for a module")
    fmt.Println("  run <id> [args...]      Run a module by full id")
    fmt.Println("  update                  Pull latest updates from Git")
    fmt.Println("  new | scaffold          Create a new module skeleton")
    fmt.Println("  version                 Show SHF version")
    fmt.Println()
    fmt.Println("Short execution via aliases:")
    fmt.Println("  shf hash <file>")
    fmt.Println("  shf sshlog <logfile>")
    fmt.Println("  shf ti-ip <ip>")
}
