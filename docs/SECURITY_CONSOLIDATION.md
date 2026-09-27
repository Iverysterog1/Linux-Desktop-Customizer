# Legacy workflow security consolidation

The repository historically accumulated multiple automated recovery/publication workflows. Several could write or force-write repository refs. The target canonical design is simpler:

1. validation workflows are read-only;
2. recovery material never grants publication authority;
3. repository changes use topic branch -> validation -> reviewed PR;
4. releases/tags are a separate explicitly authorized operation;
5. retired publication workflows remain inert/auditable until safely removed in a later cleanup.

The canonical integration branch is absorbing the individual security PRs by affected workflow. Overlapping PRs are not blindly merged because that would duplicate guards and create conflicts without increasing protection.

Feature development remains paused until every write-capable legacy workflow on the default branch has a canonical disposition.
