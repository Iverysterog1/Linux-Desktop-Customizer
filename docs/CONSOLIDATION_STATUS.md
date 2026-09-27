# Consolidation status

Date: 2026-09-27

Canonical integration branch: `integration/canonical`
Base: `main` at `56fa162d076cbe96c3ebdca828f480070293708a`

## Integrated centrally

- PR #20 family: retirement of `finalize-reviewed-source.yml`, `sync-final-source.yml`, and `tag-v0.9.0.yml`, with a regression guard.
- PR #19: reviewed GitHub Actions dependency-update configuration.
- PR #21: least-privilege pull-request guard against workflow write/direct-push regressions.
- Branch governance defining `main` as stable and `integration/canonical` as the single consolidation line.

## Preserved as recovery evidence, not canonical application source

- PR #5: preserved 0.11 source package validator; current evidence says the stream is truncated.
- PR #9: historical blob archive validator; current evidence says an authoritative blob is unavailable.

These validators and their findings must remain available, but their failed recovery inputs must not be promoted into canonical application source.

## Legacy security PR family still to absorb/classify

PRs #1, #3, #7, #8, #10-#18 address additional legacy workflows. They remain open evidence until each unique workflow remediation is represented on `integration/canonical`. Do not delete their branches before that classification is complete.

## Development state

PAUSED for feature work until consolidation completes. No merge to `main`, release, tag publication, credential rotation, or branch deletion is authorized by this document.
