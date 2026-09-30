#!/usr/bin/env bash
set -euo pipefail

BIN_DIR="${LTC_INSTALL_PREFIX:-${XDG_BIN_HOME:-$HOME/.local/bin}}"
DATA_HOME="${XDG_DATA_HOME:-$HOME/.local/share}"
DESKTOP_ID="io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

# Ordinary uninstall removes only the installation payload owned by this
# installer. Configuration, profiles, snapshots/history and other user data
# are intentionally preserved. Data cleanup/restoration is a separate action.
rm -f -- "$BIN_DIR/ltc" "$BIN_DIR/ltc-ui"
rm -f -- "$DATA_HOME/applications/$DESKTOP_ID"

echo "Removed Linux Desktop Customizer user installation"
echo "Preserved Linux Desktop Customizer user data and recovery state"
