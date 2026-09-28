#!/usr/bin/env bash
set -euo pipefail

echo 'BLOCKED: historical beta blob reconstruction has an unavailable authoritative object.'
echo 'Do not substitute guessed bytes. Recovery must establish object provenance and the pinned complete-archive digest before source promotion.'
echo 'The preserved read-only reconstruction procedure is scripts/recovery/validate-historical-beta-blobs.sh.'
exit 1
