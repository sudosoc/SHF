package core

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
)


func resultsDir(moduleID string) string {
    exe, err := os.Executable()
    if err != nil {
        return ""
    }

    // فولدر الفريم وورك نفسه
    root := filepath.Dir(exe)

    // استخراج الكاتيجوري (offensive / defensive / forensics / threat_intel)
    parts := strings.Split(moduleID, "/")
    category := parts[0]

    // مجلد النتائج داخل الفريم وورك
    results := filepath.Join(root, "results", category)

    // نضمن وجود المجلد
    os.MkdirAll(results, 0o755)

    return results
}



func saveJSONIfValid(output []byte, moduleID string) {
    var js interface{}
    if err := json.Unmarshal(output, &js); err != nil {
        // لو مش JSON.. خلاص
        return
    }

    // نحدد مجلد النتائج حسب نوع الموديول
    dir := resultsDir(moduleID)
    if dir == "" {
        return
    }

    // نحول الموديول ID لاسم ملف
    // مثال: forensics/hash_file → hash_file
    parts := strings.Split(moduleID, "/")
    name := parts[len(parts)-1]

    ts := time.Now().UTC().Format("20060102T150405Z")
    filename := fmt.Sprintf("%s_%s.json", name, ts)
    path := filepath.Join(dir, filename)

    if err := os.WriteFile(path, output, 0o644); err != nil {
        return
    }

    fmt.Println("[+] JSON result saved to:", path)
}


func RunModule(m Module, cfg Config, logger Logger, directArgs []string, smartPrimary bool) {
    cmdName := ""
    args := []string{}

    switch m.Type {
    case "python":
        cmdName = "python3"
        scriptPath := filepath.Join(m.BaseDir, m.Path)
        args = append(args, scriptPath)
    case "bash":
        cmdName = "bash"
        scriptPath := filepath.Join(m.BaseDir, m.Path)
        args = append(args, scriptPath)
    default:
        cmdName = filepath.Join(m.BaseDir, m.Path)
    }

    finalArgs := []string{}
    if smartPrimary && m.PrimaryArg != "" &&
        len(directArgs) > 0 &&
        len(directArgs[0]) > 0 &&
        directArgs[0][0] != '-' {
        // shf hash file.txt → نحول أول arg لـ --file file.txt
        finalArgs = append(finalArgs, "--"+m.PrimaryArg, directArgs[0])
        finalArgs = append(finalArgs, directArgs[1:]...)
    } else {
        finalArgs = append(finalArgs, directArgs...)
    }

    cmd := exec.Command(cmdName, append(args, finalArgs...)...)

   
    out, err := cmd.CombinedOutput()

    if len(out) > 0 {
        fmt.Print(string(out))       
        saveJSONIfValid(out, m.ID)   
    }
    if err != nil {
        fmt.Println("[!] Error running module:", err)
    }

    logger.Log(m.ID, finalArgs, err)
}
