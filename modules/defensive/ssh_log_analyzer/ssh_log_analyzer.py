#!/usr/bin/env python3
import argparse
import json
import os
import re
from collections import Counter
from datetime import datetime, timezone

def main():
    parser = argparse.ArgumentParser(description="SSH auth.log analyzer for failed login attempts.")
    parser.add_argument("--log", default="/var/log/auth.log", help="Path to SSH auth log")
    parser.add_argument("--threshold", type=int, default=5, help="Minimum failures to flag IP")
    parser.add_argument("--json", action="store_true", help="Output JSON")
    parser.add_argument("positional_log", nargs="?", help="Optional positional log path")
    args = parser.parse_args()

    log_path = args.positional_log or args.log

    if not os.path.isfile(log_path):
        print(f"[!] Log file not found: {log_path}")
        return

    failed_regex = re.compile(r"Failed password.*from ([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)")
    counter = Counter()

    with open(log_path, "r", encoding="utf-8", errors="ignore") as f:
        for line in f:
            m = failed_regex.search(line)
            if m:
                counter[m.group(1)] += 1

    suspicious = {ip: count for ip, count in counter.items() if count >= args.threshold}

    if args.json:
        data = {
            "module": "defensive/ssh_log_analyzer",
            "log": os.path.abspath(log_path),
            "threshold": args.threshold,
            "suspicious": suspicious,
            "timestamp": datetime.now(timezone.utc).isoformat(),
        }
        print(json.dumps(data, indent=3))
    else:
        print("Module: defensive/ssh_log_analyzer")
        print("Log:", os.path.abspath(log_path))
        print("Threshold:", args.threshold)
        print("Suspicious IPs:")
        for ip, count in sorted(suspicious.items(), key=lambda x: -x[1]):
            print(f"  {ip} -> {count} failed attempts")

if __name__ == "__main__":
    main()
