#!/usr/bin/env bash
set -euo pipefail

files=(
  .github/workflows/finalize-reviewed-source.yml
  .github/workflows/sync-final-source.yml
  .github/workflows/tag-v0.9.0.yml
)

for file in "${files[@]}"; do
  test -f "$file"

  # These retired workflows must remain manual and read-only.
  grep -Eq '^[[:space:]]*workflow_dispatch:' "$file"
  grep -Eq '^[[:space:]]*contents:[[:space:]]*read[[:space:]]*$' "$file"

  # Fail closed if executable YAML regains an automatic push trigger,
  # write permission, or repository/tag publication command.
  if grep -Eq '^[[:space:]]+push:[[:space:]]*$|^[[:space:]]*contents:[[:space:]]*write[[:space:]]*$|^[[:space:]]+(git[[:space:]]+)?push([[:space:]]|$)|git[[:space:]]+push([[:space:]]|$)' "$file"; then
    echo "unsafe retired workflow primitive detected: $file" >&2
    exit 1
  fi
done

echo 'PASS: retired direct-main workflows remain manual and read-only.'
