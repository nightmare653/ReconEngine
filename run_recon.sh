#!/bin/bash

if [ "$1" == "" ]; then
    echo "Usage: ./run_recon.sh -d <domain> | --list <domains.txt>"
    exit 1
fi

# Build the Go recon engine binary
echo "⚙️ Building..."
go build -o reconengine ./cmd/main.go
if [ $? -ne 0 ]; then
    echo "❌ Build failed."
    exit 1
fi

# Parse arguments
if [ "$1" == "--list" ]; then
    DOMAIN_FILE="$2"
    if [ ! -f "$DOMAIN_FILE" ]; then
        echo "❌ File not found: $DOMAIN_FILE"
        exit 1
    fi
    echo "📜 Scanning list of domains from: $DOMAIN_FILE"
    ./reconengine --list "$DOMAIN_FILE"
else
    DOMAIN="$2"
    if [ "$1" == "-d" ] && [ "$DOMAIN" != "" ]; then
        echo "🔍 Launching ReconEngine on: $DOMAIN"
        ./reconengine -d "$DOMAIN"
    else
        echo "❌ Invalid arguments. Use: ./run_recon.sh -d <domain> | --list <domains.txt>"
        exit 1
    fi
fi
