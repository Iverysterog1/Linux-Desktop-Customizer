#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
BIN_DIR="${LTC_INSTALL_PREFIX:-${XDG_BIN_HOME:-$HOME/.local/bin}}"
DATA_HOME="${XDG_DATA_HOME:-$HOME/.local/share}"
APPLICATIONS_DIR="$DATA_HOME/applications"
DESKTOP_ID="io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

ensure_directory() {
  local path="$1"
  if [[ -e "$path" ]]; then
    if [[ ! -d "$path" ]]; then
      echo "Install destination exists but is not a directory: $path" >&2
      exit 1
    fi
    return
  fi
  install -d -m 0755 "$path"
}

for name in ltc ltc-ui; do
  source_path="$SCRIPT_DIR/$name"
  if [[ ! -f "$source_path" || ! -x "$source_path" ]]; then
    echo "Missing executable package payload: $source_path" >&2
    exit 1
  fi
done
if [[ ! -f "$SCRIPT_DIR/$DESKTOP_ID" ]]; then
  echo "Missing desktop entry payload: $SCRIPT_DIR/$DESKTOP_ID" >&2
  exit 1
fi

# Preserve permissions on user-owned directories that already exist.
# New directories created by this installer use a conservative 0755 mode.
ensure_directory "$BIN_DIR"
ensure_directory "$APPLICATIONS_DIR"

for name in ltc ltc-ui; do
  install -m 0755 "$SCRIPT_DIR/$name" "$BIN_DIR/$name"
done
install -m 0644 "$SCRIPT_DIR/$DESKTOP_ID" "$APPLICATIONS_DIR/$DESKTOP_ID"

echo "Installed Linux Desktop Customizer commands to $BIN_DIR"
echo "Installed desktop entry to $APPLICATIONS_DIR/$DESKTOP_ID"
