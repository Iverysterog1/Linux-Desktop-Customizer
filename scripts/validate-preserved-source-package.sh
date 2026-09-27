#!/usr/bin/env bash
set -euo pipefail

echo 'BLOCKED: preserved 0.11 recovery stream is known to be incomplete.'
echo 'Do not promote or splice recovery streams. Recover authoritative complete bytes and re-establish strict SHA-256/XZ/source validation before promotion.'
exit 1
