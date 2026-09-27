# Consolidation summary

A single integration branch now exists to replace the previous pattern of accumulating many independent long-lived branches. Security, dependency, documentation, and recovery-validation work is being represented there without changing `main`.

The branch list will remain temporarily larger than the desired end state because deletion is intentionally deferred until every branch's unique work is accounted for. The target end state is a small set: stable `main`, canonical integration, and only short-lived active topics plus explicitly preserved historical evidence where needed.
