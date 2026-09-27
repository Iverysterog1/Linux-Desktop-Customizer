#!/usr/bin/env bash
set -euo pipefail

workflows=(
  .github/workflows/finalize-reviewed-source.yml
  .github/workflows/sync-final-source.yml
  .github/workflows/tag-v0.9.0.yml
)

for file in "${workflows[@]}"; do
  test -f "$file"
  grep -q 'workflow_dispatch:' "$file"
  grep -Eq '^[[:space:]]*contents:[[:space:]]*read[[:space:]]*$' "$file"
  if grep -Eq '^[[:space:]]*push:|contents:[[:space:]]*write|git[[:space:]]+push|--force' "$file"; then
    echo "FAIL: retired writer regained publication authority: $file" >&2
    exit 1
  fi
done

echo 'Retired main-writer guard: PASS'
