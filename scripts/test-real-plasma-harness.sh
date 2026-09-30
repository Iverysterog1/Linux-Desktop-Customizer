#!/usr/bin/env bash
set -euo pipefail

if [[ ! -f go.mod || ! -f scripts/validate-real-plasma.sh ]]; then
  echo "run from repository root" >&2
  exit 2
fi

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

mkdir -p "$tmp/bin"
state="$tmp/state"
trace="$tmp/trace"
: >"$trace"
printf '%s\n' "BreezeLight" >"$state"

cat >"$tmp/bin/kreadconfig6" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
cat "$LDC_TEST_STATE"
EOF

cat >"$tmp/bin/kwriteconfig6" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
printf 'unexpected direct kwriteconfig6 invocation: %s\n' "$*" >>"$LDC_TEST_TRACE"
exit 99
EOF

cat >"$tmp/bin/plasma-apply-colorscheme" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
printf 'plasma-apply-colorscheme %s\n' "$*" >>"$LDC_TEST_TRACE"
if [[ "${1:-}" == "--list-schemes" ]]; then
  printf '%s\n' "BreezeLight" "BreezeDark"
  exit 0
fi
echo "unexpected native mutation in harness regression test" >&2
exit 99
EOF

cat >"$tmp/bin/go" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
printf 'go %s\n' "$*" >>"$LDC_TEST_TRACE"

if [[ "${1:-}" != "run" || "${2:-}" != "./cmd/ltc" ]]; then
  echo "unexpected fake go invocation: $*" >&2
  exit 98
fi

case "${3:-}" in
  status)
    printf '%s\n' '{"version":"test","kde":{"detected":true}}'
    ;;
  preview)
    printf '%s\n' '{"supported":true,"changes":[{"adapter":"kde-config","key":"color-scheme"}]}'
    ;;
  apply)
    printf '%s\n' "$LDC_TEST_TARGET" >"$LDC_TEST_STATE"
    cat <<JSON
{
  "id": "0123456789abcdef0123456789abcdef",
  "status": "applied"
}
JSON
    ;;
  rollback)
    printf '%s\n' "BreezeLight" >"$LDC_TEST_STATE"
    cat <<JSON
{
  "id": "0123456789abcdef0123456789abcdef",
  "status": "rolled_back"
}
JSON
    ;;
  *)
    echo "unexpected fake ltc command: $*" >&2
    exit 97
    ;;
esac
EOF

chmod +x "$tmp/bin/"*

export PATH="$tmp/bin:$PATH"
export LDC_TEST_STATE="$state"
export LDC_TEST_TRACE="$trace"
export LDC_TEST_TARGET="BreezeDark"
export XDG_CURRENT_DESKTOP="KDE"
export XDG_SESSION_TYPE="wayland"

reset_fixture() {
  printf '%s\n' "BreezeLight" >"$state"
  : >"$trace"
}

assert_no_apply() {
  if grep -Fq "go run ./cmd/ltc apply --profile" "$trace"; then
    echo "unexpected mutation path reached" >&2
    cat "$trace" >&2
    exit 1
  fi
}

reset_fixture
read_only_out="$tmp/read-only.out"
bash scripts/validate-real-plasma.sh >"$read_only_out"
grep -Fq "MUTATION_RESULT=NOT RUN" "$read_only_out"
grep -Fq "current_color_scheme=BreezeLight" "$read_only_out"
assert_no_apply

reset_fixture
blocked_out="$tmp/blocked.out"
set +e
bash scripts/validate-real-plasma.sh --apply --target BreezeDark >"$blocked_out" 2>&1
blocked_status=$?
set -e
[[ "$blocked_status" -eq 2 ]]
grep -Fq "RESULT=BLOCKED" "$blocked_out"
grep -Fq "set LDC_REAL_PLASMA_MUTATION=YES" "$blocked_out"
assert_no_apply

reset_fixture
invalid_out="$tmp/invalid.out"
set +e
LDC_REAL_PLASMA_MUTATION=YES bash scripts/validate-real-plasma.sh --apply --target . >"$invalid_out" 2>&1
invalid_status=$?
set -e
[[ "$invalid_status" -eq 2 ]]
grep -Fq "RESULT=BLOCKED" "$invalid_out"
grep -Fq "simple installed scheme identifier" "$invalid_out"
assert_no_apply

reset_fixture
pass_out="$tmp/pass.out"
LDC_REAL_PLASMA_MUTATION=YES \
LDC_REAL_PLASMA_VISIBLE_RESULT=PASS \
  bash scripts/validate-real-plasma.sh --apply --target BreezeDark >"$pass_out"
grep -Fq "CONFIG_APPLY=PASS" "$pass_out"
grep -Fq "VISIBLE_APPLY=PASS" "$pass_out"
grep -Fq "ROLLBACK=PASS" "$pass_out"
grep -Fq "RESULT=PASS" "$pass_out"
[[ "$(cat "$state")" == "BreezeLight" ]]
grep -Fq "go run ./cmd/ltc apply --profile" "$trace"
grep -Fq "go run ./cmd/ltc rollback --transaction 0123456789abcdef0123456789abcdef" "$trace"

reset_fixture
fail_out="$tmp/fail.out"
set +e
LDC_REAL_PLASMA_MUTATION=YES \
LDC_REAL_PLASMA_VISIBLE_RESULT=FAIL \
  bash scripts/validate-real-plasma.sh --apply --target BreezeDark >"$fail_out" 2>&1
fail_status=$?
set -e
[[ "$fail_status" -eq 1 ]]
grep -Fq "CONFIG_APPLY=PASS" "$fail_out"
grep -Fq "VISIBLE_APPLY=FAIL" "$fail_out"
grep -Fq "ROLLBACK=PASS" "$fail_out"
grep -Fq "RESULT=FAIL" "$fail_out"
[[ "$(cat "$state")" == "BreezeLight" ]]

echo "real Plasma harness regression checks: PASS"
