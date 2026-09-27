# Desired repository end state

After consolidation and later explicit cleanup:

- `main`: stable reviewed source/release line;
- `integration/canonical`: current integrated development line;
- short-lived feature/fix branches only while actively reviewed;
- historical/recovery evidence preserved by commits, PRs, tags/artifacts or clearly named archival refs only when genuinely necessary;
- no fleet of overlapping security branches representing the same policy change.

This end state is reached through evidence-preserving cleanup, not destructive ref rewriting.
