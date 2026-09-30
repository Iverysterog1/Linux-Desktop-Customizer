#!/usr/bin/env bash
set -euo pipefail

ROOT="$(mktemp -d)"
trap 'rm -rf "$ROOT"' EXIT

PACKAGE="$ROOT/package"
HOME_DIR="$ROOT/home"
DATA_HOME="$ROOT/data"
CONFIG_HOME="$ROOT/config"
STATE_HOME="$ROOT/state"
BIN_DIR="$HOME_DIR/.local/bin"
APPLICATIONS_DIR="$DATA_HOME/applications"

mkdir -p "$PACKAGE" "$HOME_DIR" "$DATA_HOME" "$CONFIG_HOME" "$STATE_HOME"

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

mkdir -p "$BIN_DIR" "$APPLICATIONS_DIR"
chmod 0700 "$BIN_DIR" "$APPLICATIONS_DIR"

mkdir -p "$CONFIG_HOME/linux-desktop-customizer"          "$STATE_HOME/linux-desktop-customizer/history"          "$DATA_HOME/linux-desktop-customizer"
printf '%s
' '{"name":"keep-me"}' > "$CONFIG_HOME/linux-desktop-customizer/profile.json"
printf '%s
' 'snapshot' > "$STATE_HOME/linux-desktop-customizer/history/snapshot.txt"
printf '%s
' 'user-theme' > "$DATA_HOME/linux-desktop-customizer/user-theme.txt"

run_with_env() {
  HOME="$HOME_DIR"   XDG_DATA_HOME="$DATA_HOME"   XDG_CONFIG_HOME="$CONFIG_HOME"   XDG_STATE_HOME="$STATE_HOME"   "$@"
}

run_with_env "$PACKAGE/install-user.sh"

test -x "$BIN_DIR/ltc"
test -x "$BIN_DIR/ltc-ui"
test "$("$BIN_DIR/ltc")" = "ltc-smoke"
test "$("$BIN_DIR/ltc-ui")" = "ltc-ui-smoke"
test -f "$APPLICATIONS_DIR/io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

test "$(stat -c '%a' "$BIN_DIR")" = "700"
test "$(stat -c '%a' "$APPLICATIONS_DIR")" = "700"
test "$(stat -c '%a' "$BIN_DIR/ltc")" = "755"
test "$(stat -c '%a' "$BIN_DIR/ltc-ui")" = "755"
test "$(stat -c '%a' "$APPLICATIONS_DIR/io.github.iverysterog1.LinuxDesktopCustomizer.desktop")" = "644"

run_with_env "$PACKAGE/uninstall-user.sh"

test ! -e "$BIN_DIR/ltc"
test ! -e "$BIN_DIR/ltc-ui"
test ! -e "$APPLICATIONS_DIR/io.github.iverysterog1.LinuxDesktopCustomizer.desktop"

test -f "$CONFIG_HOME/linux-desktop-customizer/profile.json"
test -f "$STATE_HOME/linux-desktop-customizer/history/snapshot.txt"
test -f "$DATA_HOME/linux-desktop-customizer/user-theme.txt"
test "$(stat -c '%a' "$BIN_DIR")" = "700"
test "$(stat -c '%a' "$APPLICATIONS_DIR")" = "700"

run_with_env "$PACKAGE/uninstall-user.sh"

echo "PASS: clean user install/uninstall + permissions/data-preservation smoke test"
