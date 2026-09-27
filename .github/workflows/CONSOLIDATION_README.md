# Workflow consolidation

Legacy publication/recovery workflows in this directory are being retired into read-only validation or inert audit tombstones. New workflows must use least privilege and must not push directly or forcibly to `main`.

The changed-workflow regression guard on `integration/canonical` is the central policy check. Do not create additional one-off publication workflows to bypass it.
