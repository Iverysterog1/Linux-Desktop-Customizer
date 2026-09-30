#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'EOF'
Real KDE Plasma validation harness.

Read-only by default:
  scripts/validate-real-plasma.sh

Explicit mutation test:
  LDC_REAL_PLASMA_MUTATION=YES scripts/validate-real-plasma.sh --apply --target BreezeDark

Optional non-interactive visible result:
  LDC_REAL_PLASMA_VISIBLE_RESULT=PASS|FAIL

The harness always attempts rollback after an apply. It does not publish evidence.
EOF
}

apply=false
target=""
while (($#)); do
  case "$1" in
    --apply) apply=true ;;
    --target)
      shift
      target="${1:-}"
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "unknown argument: $1" >&2
      usage >&2
      exit 2
      ;;
  esac
  shift
done

if [[ ! -f go.mod ]]; then
  echo "BLOCKED: run from the repository root (go.mod not found)" >&2
  exit 2
fi

desktop="${XDG_CURRENT_DESKTOP:-}"
session_type="${XDG_SESSION_TYPE:-unknown}"
if [[ "${desktop,,}" != *kde* && "${desktop,,}" != *plasma* && -z "${KDE_FULL_SESSION:-}" ]]; then
  echo "RESULT=BLOCKED"
  echo "reason=KDE Plasma session not detected"
  exit 2
fi

for command in kreadconfig6 kwriteconfig6 plasma-apply-colorscheme go; do
  if ! command -v "$command" >/dev/null 2>&1; then
    echo "RESULT=BLOCKED"
    echo "reason=$command not found in PATH"
    exit 2
  fi
done

sentinel="__LDC_REAL_PLASMA_MISSING__"
read_scheme() {
  kreadconfig6 --file kdeglobals --group General --key ColorScheme --default "$sentinel"
}

before="$(read_scheme)"

echo "phase=preflight"
echo "desktop=$desktop"
echo "session_type=$session_type"
echo "current_color_scheme=$before"
echo "native_apply_tool=$(command -v plasma-apply-colorscheme)"
echo "ltc_status_begin"
go run ./cmd/ltc status
echo "ltc_status_end"
echo "native_scheme_list_begin"
LC_ALL=C plasma-apply-colorscheme --list-schemes || true
echo "native_scheme_list_end"

if [[ "$apply" != true ]]; then
  echo "MUTATION_RESULT=NOT RUN"
  echo "reason=read-only preflight complete; rerun with explicit mutation opt-in to test apply/rollback"
  exit 0
fi

if [[ "${LDC_REAL_PLASMA_MUTATION:-}" != "YES" ]]; then
  echo "RESULT=BLOCKED"
  echo "reason=set LDC_REAL_PLASMA_MUTATION=YES for the explicit mutation test"
  exit 2
fi

if [[ -z "$target" ]]; then
  echo "RESULT=BLOCKED"
  echo "reason=--target is required for mutation"
  exit 2
fi

if [[ "$target" == "." || "$target" == ".." || ! "$target" =~ ^[A-Za-z0-9._+-]+$ ]]; then
  echo "RESULT=BLOCKED"
  echo "reason=validation target must be a simple installed scheme identifier"
  exit 2
fi

if [[ "$target" == "$before" ]]; then
  echo "RESULT=BLOCKED"
  echo "reason=target scheme equals current scheme; choose a visibly different installed scheme"
  exit 2
fi

profile="$(mktemp)"
apply_json="$(mktemp)"
txid=""
rollback_done=false

cleanup() {
  if [[ -n "$txid" && "$rollback_done" != true ]]; then
    rollback_done=true
    echo "emergency_rollback=attempt"
    go run ./cmd/ltc rollback --transaction "$txid" >/dev/null 2>&1 || true
  fi
  rm -f "$profile" "$apply_json"
}
trap cleanup EXIT
trap 'exit 130' INT TERM

cat >"$profile" <<EOF
{
  "version": 1,
  "name": "Real Plasma color-scheme validation",
  "operations": [
    {
      "adapter": "kde-config",
      "action": "set",
      "key": "color-scheme",
      "value": "$target"
    }
  ]
}
EOF

echo "phase=preview"
go run ./cmd/ltc preview --profile "$profile"

echo "phase=apply"
go run ./cmd/ltc apply --profile "$profile" | tee "$apply_json"
txid="$(sed -n 's/^[[:space:]]*"id":[[:space:]]*"\([^"]*\)".*/\1/p' "$apply_json" | head -n1)"
if [[ -z "$txid" ]]; then
  echo "RESULT=FAIL"
  echo "reason=apply returned no transaction id"
  exit 1
fi

after_apply="$(read_scheme)"
config_apply="FAIL"
if [[ "$after_apply" == "$target" ]]; then
  config_apply="PASS"
fi

visible="${LDC_REAL_PLASMA_VISIBLE_RESULT:-}"
if [[ -z "$visible" && -t 0 ]]; then
  echo
  echo "MANUAL VISIBLE CHECK:"
  echo "Confirm whether Plasma/application colors visibly changed to the target scheme."
  echo "Type PASS only if the change is genuinely visible; otherwise type FAIL."
  read -r visible
fi
case "$visible" in
  PASS|FAIL) ;;
  *) visible="NOT RUN" ;;
esac

echo "phase=rollback"
go run ./cmd/ltc rollback --transaction "$txid"
rollback_done=true

after_rollback="$(read_scheme)"
rollback_result="FAIL"
if [[ "$after_rollback" == "$before" ]]; then
  rollback_result="PASS"
fi

echo "CONFIG_APPLY=$config_apply"
echo "VISIBLE_APPLY=$visible"
echo "ROLLBACK=$rollback_result"
echo "before=$before"
echo "after_apply=$after_apply"
echo "after_rollback=$after_rollback"
echo "transaction_id=$txid"

if [[ "$config_apply" == PASS && "$visible" == PASS && "$rollback_result" == PASS ]]; then
  echo "RESULT=PASS"
  exit 0
fi

if [[ "$visible" == "NOT RUN" && "$config_apply" == PASS && "$rollback_result" == PASS ]]; then
  echo "RESULT=NOT RUN"
  echo "reason=config apply/rollback succeeded but visible-effect confirmation was not performed"
  exit 0
fi

echo "RESULT=FAIL"
exit 1
