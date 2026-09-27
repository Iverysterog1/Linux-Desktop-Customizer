#!/usr/bin/env bash
set -euo pipefail

files=("$@")
if [ ${#files[@]} -eq 0 ]; then
  echo 'No changed workflow files to inspect.'
  exit 0
fi

failed=0
for file in "${files[@]}"; do
  [[ "$file" == .github/workflows/*.yml || "$file" == .github/workflows/*.yaml ]] || continue
  [[ -f "$file" ]] || continue
  echo "Auditing $file"
  if grep -En '^[[:space:]]*(contents|actions|packages|id-token):[[:space:]]*write([[:space:]]|$)' "$file"; then failed=1; fi
  if grep -En "git[[:space:]]+push|--force([^[:alnum:]-]|$)|:[[:space:]]*main([[:space:]\"']|$)" "$file"; then failed=1; fi
  if grep -En '^[[:space:]]*(GH_TOKEN|GITHUB_TOKEN):' "$file"; then failed=1; fi
done

if [ "$failed" -ne 0 ]; then
  echo 'Unsafe workflow write/publication primitive detected.' >&2
  exit 1
fi

echo 'Workflow write regression guard: PASS'
