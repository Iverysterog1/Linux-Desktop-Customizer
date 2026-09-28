#!/usr/bin/env bash
set -euo pipefail

file="packaging/io.github.iverysterog1.LinuxDesktopCustomizer.desktop"
test -f "$file"

grep -Fxq '[Desktop Entry]' "$file"
grep -Fxq 'Type=Application' "$file"
grep -Fxq 'Name=Linux Desktop Customizer' "$file"
grep -Fxq 'Exec=ltc-ui' "$file"
grep -Fxq 'Terminal=false' "$file"

if command -v desktop-file-validate >/dev/null 2>&1; then
  desktop-file-validate "$file"
fi

echo "PASS: desktop entry validation"
