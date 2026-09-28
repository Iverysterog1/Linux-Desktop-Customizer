#!/usr/bin/env bash
set -euo pipefail

ROOT="$(mktemp -d)"
trap 'rm -rf "$ROOT"' EXIT

PACKAGE="$ROOT/package"
HOME_DIR="$ROOT/home"
DATA_HOME="$ROOT/data"
mkdir -p "$PACKAGE" "$HOME_DIR" "$DATA_HOME"

cp packaging/install-user.sh packaging/uninstall-user.sh packaging/io.github.iverysterog1.LinuxDesktopCustomizer.desktop "$PACKAGE/"
cat > "$PACKAGE/ltc" <<'EOF'
#!/usr/bin/env sh
echo ltc-smoke
EOF
cat > "$PACKAGE/ltc-ui" <<'EOF'
#!/usr/bin/env sh
echo ltc-ui-smoke
EOF
chmod 0755 "$PACKAGE/ltc" "$PACKAGE/ltc-ui" "$PACKAGE/install-user.sh" "$PACKAGE/uninstall-user.sh"

HOME="$HOME_DIR" XDG_DATA_HOME="$DATA_HOME" "$PACKAGE/install-user.sh"
test -x "$HOME_DIR/.local/bin/ltc"
test -x "$HOME_DIR/.local/bin/ltc-ui"
test "$("$HOME_DIR/.local/bin/ltc")" = "ltc-smoke"
test "$("$HOME_DIR/.local/bin/ltc-ui")" = "ltc-ui-smoke"
test -f "$DATA_HOME/applications/io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

HOME="$HOME_DIR" XDG_DATA_HOME="$DATA_HOME" "$PACKAGE/uninstall-user.sh"
test ! -e "$HOME_DIR/.local/bin/ltc"
test ! -e "$HOME_DIR/.local/bin/ltc-ui"
test ! -e "$DATA_HOME/applications/io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

HOME="$HOME_DIR" XDG_DATA_HOME="$DATA_HOME" "$PACKAGE/uninstall-user.sh"

echo "PASS: clean user install/uninstall + desktop entry smoke test"
