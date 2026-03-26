#!/usr/bin/env bash

echo "=== SYSTEM INFO ==="

echo "User: $(whoami)"
echo "Hostname: $(hostname)"
echo "OS: $(uname -s)"
echo "Kernel: $(uname -r)"
echo "Architecture: $(uname -m)"

echo ""
echo "=== RESOURCES ==="

echo "CPU cores: $(nproc)"
echo "Memory:"
free -h

echo ""
echo "Disk usage:"
df -h /

echo ""
echo "Uptime:"
uptime -p
