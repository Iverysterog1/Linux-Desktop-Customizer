#!/usr/bin/env bash
set -euo pipefail

PREFIX="${LTC_INSTALL_PREFIX:-${XDG_BIN_HOME:-$HOME/.local/bin}}"

rm -f -- "$PREFIX/ltc" "$PREFIX/ltc-ui"
echo "Removed Linux Desktop Customizer commands from $PREFIX"
