# Consolidation inventory

This inventory prevents branch count from being mistaken for canonical state.

## Canonical

- `main`: stable default branch; no direct writes.
- `integration/canonical`: sole integration line during consolidation.

## Historical product/recovery branches

- `effects-composer-beta`
- `effects-import`
- `publish-0.11-effects-composer`

Preserve until source provenance/recovery decisions are complete. They are not automatically canonical merely because they contain newer-looking material.

## Documentation / quality branches

- `docs/github-page-refresh`
- `quality/validate-preserved-source-package`
- `quality/validate-historical-beta-blob-archive`

Their unique governance/status/recovery findings are represented on the canonical integration line. Original branches remain evidence until PR cleanup is explicitly performed.

## Security branches

The `security/*` branches are remediation/evidence branches. Their changes are being absorbed into the canonical integration line by objective rather than by blindly merging every branch. This avoids duplicated guards and conflicting edits to the same legacy workflow.

## Cleanup rule

A branch is eligible for deletion only when its unique work is confirmed integrated or intentionally superseded and its PR/history remains sufficient for audit. Consolidation itself does not authorize deletion or merge to `main`.
