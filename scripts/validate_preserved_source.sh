#!/usr/bin/env bash
set -Eeuo pipefail

readonly SOURCE_COMMIT="${SOURCE_COMMIT:-80270fe4e7d9f01f2e4438e9d7c2bc1658088b8e}"
readonly PART_COUNT=10
readonly WORK_DIR="${WORK_DIR:-$(mktemp -d)}"
readonly ARCHIVE="$WORK_DIR/source.tar.xz"
readonly EXTRACT_DIR="$WORK_DIR/source"

cleanup() {
  if [[ "${KEEP_WORK_DIR:-0}" != "1" ]]; then
    rm -rf "$WORK_DIR"
  fi
}
trap cleanup EXIT

command -v git >/dev/null
command -v base64 >/dev/null
command -v xz >/dev/null
command -v tar >/dev/null
command -v sha256sum >/dev/null

git cat-file -e "${SOURCE_COMMIT}^{commit}"

: > "$WORK_DIR/source.b64"
for i in $(seq -w 0 $((PART_COUNT - 1))); do
  path="source-package/part${i}.b64"
  git cat-file -e "${SOURCE_COMMIT}:${path}"
  bytes=$(git cat-file -s "${SOURCE_COMMIT}:${path}")
  if [[ "$bytes" -ne 16000 ]]; then
    echo "ERROR: ${path} is ${bytes} bytes; expected 16000" >&2
    exit 1
  fi
  git show "${SOURCE_COMMIT}:${path}" >> "$WORK_DIR/source.b64"
done

base64 --decode "$WORK_DIR/source.b64" > "$ARCHIVE"
archive_bytes=$(wc -c < "$ARCHIVE" | tr -d ' ')
archive_sha256=$(sha256sum "$ARCHIVE" | awk '{print $1}')
printf 'archive_bytes=%s\narchive_sha256=%s\n' "$archive_bytes" "$archive_sha256"

# Integrity failure is fatal. Never mask xz errors with `|| true`.
xz -t "$ARCHIVE"

tar -tJf "$ARCHIVE" > "$WORK_DIR/archive-files.txt"
if grep -Eq '(^|/)\.\.(/|$)|^/' "$WORK_DIR/archive-files.txt"; then
  echo "ERROR: archive contains unsafe absolute or parent-traversal paths" >&2
  exit 1
fi

mkdir -p "$EXTRACT_DIR"
tar -xJf "$ARCHIVE" -C "$EXTRACT_DIR"

# Archives may contain a single top-level directory. Resolve the source root
# by locating VERSION rather than assuming an archive layout.
mapfile -t version_files < <(find "$EXTRACT_DIR" -type f -name VERSION -print)
if [[ "${#version_files[@]}" -ne 1 ]]; then
  echo "ERROR: expected exactly one VERSION file; found ${#version_files[@]}" >&2
  exit 1
fi
source_root=$(dirname "${version_files[0]}")
version=$(tr -d '\r\n' < "${version_files[0]}")
printf 'source_root=%s\nversion=%s\n' "$source_root" "$version"

required=(LICENSE README.md go.mod)
for path in "${required[@]}"; do
  if [[ ! -f "$source_root/$path" ]]; then
    echo "ERROR: required source file missing: $path" >&2
    exit 1
  fi
done

printf '\nRelease-readiness inventory:\n'
for pattern in '*.desktop' '*metainfo*.xml' '*appdata*.xml' '*.png' '*.svg' 'install*.sh' 'uninstall*.sh'; do
  count=$(find "$source_root" -type f -name "$pattern" | wc -l | tr -d ' ')
  printf '%-20s %s\n' "$pattern" "$count"
done

if command -v go >/dev/null; then
  (
    cd "$source_root"
    go test ./...
    go vet ./...
    go build ./cmd/ltc
    go build ./cmd/ltc-ui
  )
else
  echo "ERROR: Go is required for source validation" >&2
  exit 1
fi

printf '\nPASS: preserved source package is structurally valid and builds.\n'
