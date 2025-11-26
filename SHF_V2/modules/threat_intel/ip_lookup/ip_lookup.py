#!/usr/bin/env python3
import argparse
import json
from datetime import datetime, timezone

DEMO_REPUTATION = {
    "185.220.101.1": "Known TOR exit node (demo only)",
    "203.0.113.5": "Example scanner IP (demo only)",
}

def main():
    parser = argparse.ArgumentParser(description="Local IP reputation demo.")
    parser.add_argument("--ip", help="IP address to check")
    parser.add_argument("--json", action="store_true", help="Output JSON")
    parser.add_argument("positional_ip", nargs="?", help="Optional positional IP argument")
    args = parser.parse_args()

    ip = args.ip or args.positional_ip
    if not ip:
        print("[!] No IP specified. Use --ip or positional argument.")
        return

    info = DEMO_REPUTATION.get(ip)
    reputation = "suspicious" if info else "unknown"

    if args.json:
        data = {
            "module": "threat_intel/ip_lookup",
            "ip": ip,
            "reputation": reputation,
            "details": info or "",
            "timestamp": datetime.now(timezone.utc).isoformat(),
        }
        print(json.dumps(data, indent=3))
    else:
        print("Module: threat_intel/ip_lookup")
        print("IP:", ip)
        print("Reputation:", reputation)
        if info:
            print("Details:", info)

if __name__ == "__main__":
    main()
