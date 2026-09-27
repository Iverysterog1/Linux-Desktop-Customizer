#!/usr/bin/env bash
set -euo pipefail

base_ref="${1:-origin/main}"

mapfile -t workflows < <(git diff --name-only --diff-filter=ACMR "$base_ref"...HEAD -- '.github/workflows/*.yml' '.github/workflows/*.yaml')

if ((${#workflows[@]} == 0)); then
  echo "No changed GitHub Actions workflows to inspect."
  exit 0
fi

failed=0
for file in "${workflows[@]}"; do
  [[ -f "$file" ]] || continue
  echo "Inspecting $file"

  # Changed workflows must remain least-privilege and must not regain repository
  # publication primitives. Release/publishing authority belongs in a separately
  # reviewed design, never as an incidental workflow change.
  if grep -En '^[[:space:]]*(contents|actions|packages|id-token):[[:space:]]*write([[:space:]]*(#.*)?)?$' "$file"; then
    echo "ERROR: write-capable GITHUB_TOKEN permission in $file" >&2
    failed=1
  fi
  if grep -En '(^|[[:space:];|&])(git[[:space:]]+push)([[:space:]]|$)' "$file"; then
    echo "ERROR: git push publication primitive in $file" >&2
    failed=1
  fi
  if grep -En -- '(^|[[:space:]])--force(-with-lease)?([[:space:]]|$)|refs/heads/main|:[[:space:]]*main([[:space:]]|$)' "$file"; then
    echo "ERROR: force/default-branch publication primitive in $file" >&2
    failed=1
  fi
  if grep -En '(GH_TOKEN|GITHUB_TOKEN):[[:space:]]*\$\{\{' "$file"; then
    echo "ERROR: workflow token exported into a step environment in $file" >&2
    failed=1
  fi
done

if ((failed)); then
  echo "Workflow security regression check FAILED." >&2
  exit 1
fi

echo "Workflow security regression check PASS."
