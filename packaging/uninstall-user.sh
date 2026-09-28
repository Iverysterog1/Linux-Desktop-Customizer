#!/usr/bin/env bash
set -euo pipefail

BIN_DIR="${LTC_INSTALL_PREFIX:-${XDG_BIN_HOME:-$HOME/.local/bin}}"
DATA_HOME="${XDG_DATA_HOME:-$HOME/.local/share}"
DESKTOP_ID="io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

rm -f -- "$BIN_DIR/ltc" "$BIN_DIR/ltc-ui"
rm -f -- "$DATA_HOME/applications/$DESKTOP_ID"

echo "Removed Linux Desktop Customizer user installation"
