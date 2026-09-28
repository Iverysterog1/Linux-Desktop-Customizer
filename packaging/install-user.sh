#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
PREFIX="${LTC_INSTALL_PREFIX:-${XDG_BIN_HOME:-$HOME/.local/bin}}"

for name in ltc ltc-ui; do
  source_path="$SCRIPT_DIR/$name"
  if [[ ! -f "$source_path" || ! -x "$source_path" ]]; then
    echo "Missing executable package payload: $source_path" >&2
    exit 1
  fi
done

mkdir -p "$PREFIX"
chmod 0755 "$PREFIX"

for name in ltc ltc-ui; do
  install -m 0755 "$SCRIPT_DIR/$name" "$PREFIX/$name"
done

echo "Installed Linux Desktop Customizer commands to $PREFIX"
