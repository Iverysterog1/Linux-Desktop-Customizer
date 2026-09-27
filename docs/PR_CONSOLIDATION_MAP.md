# Pull-request consolidation map

## Absorbed into `integration/canonical`

- #4: canonical status/release governance subset.
- #5: recovery blocker/disposition preserved fail-closed; incomplete 0.11 material is not promoted.
- #9: historical recovery blocker/disposition preserved fail-closed; unavailable authoritative object is not substituted.
- #19: reviewed GitHub Actions dependency-update configuration.
- #20: three legacy main/tag writer retirements plus regression protection.
- #21: repository-wide changed-workflow write/direct-push regression guard.

## Security PRs requiring per-workflow absorption before cleanup

- #1, #3, #7, #8, #10, #11, #12, #13, #14, #15, #16, #17, #18.

Some overlap with #20. Where a workflow is already represented by #20 on the canonical line, the older PR is superseded for integration purposes but remains audit history until deliberately closed/cleaned. Unique workflow remediations must still be copied/validated before being classified as absorbed.

## No destructive cleanup yet

PR closure and branch deletion are deliberately deferred until the canonical line contains every unique remediation and the resulting diff/CI has been reviewed. This prevents branch-count cleanup from causing loss of work.
