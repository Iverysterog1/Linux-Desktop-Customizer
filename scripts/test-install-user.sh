#!/usr/bin/env bash
set -euo pipefail

ROOT="$(mktemp -d)"
trap 'rm -rf "$ROOT"' EXIT

PACKAGE="$ROOT/package"
HOME_DIR="$ROOT/home"
mkdir -p "$PACKAGE" "$HOME_DIR"

cp packaging/install-user.sh packaging/uninstall-user.sh "$PACKAGE/"
cat > "$PACKAGE/ltc" <<'EOF'
#!/usr/bin/env sh
echo ltc-smoke
EOF
cat > "$PACKAGE/ltc-ui" <<'EOF'
#!/usr/bin/env sh
echo ltc-ui-smoke
EOF
chmod 0755 "$PACKAGE/ltc" "$PACKAGE/ltc-ui" "$PACKAGE/install-user.sh" "$PACKAGE/uninstall-user.sh"

HOME="$HOME_DIR" "$PACKAGE/install-user.sh"
test -x "$HOME_DIR/.local/bin/ltc"
test -x "$HOME_DIR/.local/bin/ltc-ui"
test "$("$HOME_DIR/.local/bin/ltc")" = "ltc-smoke"
test "$("$HOME_DIR/.local/bin/ltc-ui")" = "ltc-ui-smoke"

HOME="$HOME_DIR" "$PACKAGE/uninstall-user.sh"
test ! -e "$HOME_DIR/.local/bin/ltc"
test ! -e "$HOME_DIR/.local/bin/ltc-ui"

HOME="$HOME_DIR" "$PACKAGE/uninstall-user.sh"

echo "PASS: clean user install/uninstall smoke test"
