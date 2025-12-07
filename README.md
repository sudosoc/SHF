# SHF - SudoSoc Hybrid Framework  

[![License](https://img.shields.io/badge/License-Proprietary-red.svg)](./LICENSE)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20Windows-blue.svg)
![Language](https://img.shields.io/badge/Core-Go%20%7C%20Rust%20%7C%20Python-green.svg)
![Status](https://img.shields.io/badge/Status-Active-brightgreen.svg)

SHF (sudosoc Hybrid Framework) is a modular cyber security framework that integrates:

- 🔴 **Offensive modules** (Red Team)
- 🔵 **Defensive modules** (Blue Team)
- 🟣 **DFIR / Forensics**
- 🟡 **Threat Intelligence**

The framework is built for performance, modularity, and easy extensibility —  
allowing you to plug in new tools and scripts effortlessly using the SHF Dev Studio.

> ⚠️ This project is **proprietary** and protected under the SHF License.  
> See: [LICENSE](./LICENSE)

---

# 📦 Installation

You can install SHF using one of the following methods.

---

## 🛠️ **Linux Installation (Recommended)**

```bash
git clone https://github.com/sudosoc/SHF.git
cd SHF
chmod +x install.sh
sudo ./install.sh
```
# After installation

Run SHF globally:
```bash
shf
```
To update the framework:
```bash
shf update
```

🚀 Usage
```bash
shf
```

```bash
====================================================
  SHF - SudoSoc Hybrid Framework  (v0.x.x)
====================================================

Hybrid cyber security framework combining:
  - Offensive (Red)
  - Defensive (Blue)
  - Forensics (DFIR)
  - Threat Intelligence (TI)

Usage:
  shf [command] [options]

Commands:
  list        List all modules
  run         Run a specific module
  help        Show global help
  version     Display SHF version
```

📚 Examples
```bash
shf list
```
```bash
shf run offensive/network/port_scanner --ip 192.168.1.1 --json
```
```bash
shf run forensics/files/hash_checker -h
```

🧩 Project Structure
```bash
SHF/
│
├── cli/                     → SHF main CLI entrypoint
├── modules/                 → All offensive/defensive/forensics/TI modules
│     ├── offensive/
│     ├── defensive/
│     ├── forensics/
│     └── threat_intelligence/
│
├── internal/                → Core engine & dispatcher
├── config/                  → YAML configs
├── docs/                    → Documentation
└── install.sh               → Installer
```



🔐 License (Proprietary)

This project uses a custom proprietary license.
Redistribution, modification, or commercial usage is strictly prohibited.

See: [LICENSE](./LICENSE)



🛡️ Security Policy

Use SHF only on:

systems you own

or systems you have explicit written authorization to test

Unauthorized use is strictly prohibited.
See: [SECURITY](./SECURITY)
See: [LICENSE](./LICENSE)



👤 Author

sudosoc - Seif

🌐 Website: https://sudosoc.com

🐙 GitHub: https://github.com/sudosoc


⭐ Support & Contact


For licensing or business inquiries:

contact@sudosoc.com
