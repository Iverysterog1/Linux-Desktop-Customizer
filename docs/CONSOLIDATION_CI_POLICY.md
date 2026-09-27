# CI policy

New consolidation/security CI must use least privilege, bounded timeouts, and non-persisted checkout credentials where possible. A workflow that has not executed against the relevant canonical commit is NOT RUN, not PASS.
