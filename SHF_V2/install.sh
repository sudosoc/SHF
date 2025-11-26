#!/usr/bin/env bash
set -e

echo "============================================"
echo " Installing SHF - sudosoc Hybrid Framework "
echo "============================================"
echo " >>> To Contact us [ ceo@sudosoc.com ] <<< "
echo ""

if ! command -v go >/dev/null 2>&1; then
  echo "[!] Go is not installed. Please install Go (1.21+) and retry. >>> [sudo apt install -y golang] to install go language "
  exit 1
fi

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

echo "[+] Running go mod tidy..."
go mod tidy

echo "[+] Building SHF binary..."
go build -o shf ./cmd/shf

echo "[+] Making binary executable..."
chmod +x shf

CONFIG_DIR="$HOME/.shf"
if [ ! -d "$CONFIG_DIR" ]; then
  echo "[+] Creating config directory at $CONFIG_DIR"
  mkdir -p "$CONFIG_DIR"
fi

if [ ! -f "$CONFIG_DIR/config.yaml" ]; then
  echo "[+] Installing default config to $CONFIG_DIR/config.yaml"
  cp "$ROOT_DIR/config/default_config.yaml" "$CONFIG_DIR/config.yaml"
fi

LOG_DIR="$CONFIG_DIR/logs"
if [ ! -d "$LOG_DIR" ]; then
  echo "[+] Creating log directory at $LOG_DIR"
  mkdir -p "$LOG_DIR"
fi

RESULT_DIR="$CONFIG_DIR/results"
if [ ! -d "$RESULT_DIR" ]; then
  echo "[+] Creating results directory at $RESULT_DIR"
  mkdir -p "$RESULT_DIR"
fi


echo "[+] Attempting to create /usr/local/bin/shf symlink (requires sudo)..."
if sudo ln -sf "$ROOT_DIR/shf" /usr/local/bin/shf; then
  echo "[+] Symlink created. You can now run SHF using: shf from every where ..."
else
  echo "[!] Could not create symlink. You can still run SHF using:"
  echo "    $ROOT_DIR/shf"
fi

echo ""
echo "Done."
