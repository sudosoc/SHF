#!/usr/bin/env python3
import argparse
import hashlib
import json
import os
import sys
from datetime import datetime, timezone

def compute_sha256(path):
    h = hashlib.sha256()
    with open(path, "rb") as f:
        for chunk in iter(lambda: f.read(8192), b""):
            if not chunk:
                break
            h.update(chunk)
    return h.hexdigest()

def main():
    parser = argparse.ArgumentParser(description="Safe file hash calculator (SHA256 only).")
    parser.add_argument("--file", help="Path to file to hash")
    parser.add_argument("--algo", default="sha256", choices=["sha256"], help="Hash algorithm (only sha256 supported in this demo)")
    parser.add_argument("--json", action="store_true", help="Output JSON")
    parser.add_argument("positional_file", nargs="?", help="Optional positional file argument")
    args = parser.parse_args()

    file_path = args.file or args.positional_file
    if not file_path:
        print("[!] No file specified. Use --file or positional argument.")
        sys.exit(1)

    if not os.path.isfile(file_path):
        print(f"[!] File not found: {file_path}")
        sys.exit(1)

    try:
        digest = compute_sha256(file_path)
    except Exception as e:
        print(f"[!] Error reading file: {e}")
        sys.exit(1)

    if args.json:
        data = {
            "module": "forensics/hash_file",
            "file": os.path.abspath(file_path),
            "algo": "sha256",
            "hash": digest,
            "timestamp": datetime.now(timezone.utc).isoformat(),
        }
        print(json.dumps(data, indent=3))
    else:
        print("Module: forensics/hash_file")
        print("File:", os.path.abspath(file_path))
        print("Algorithm: sha256")
        print("Hash:", digest)

if __name__ == "__main__":
    main()
