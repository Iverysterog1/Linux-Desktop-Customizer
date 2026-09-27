# Legacy workflow retirement plan

Every legacy workflow receives one of three dispositions:

- **retain-read-only**: useful validation remains but all publication authority is removed;
- **retired-tombstone**: obsolete publication/recovery path becomes manual/read-only and documents the replacement path;
- **remove-later**: after canonical consolidation and audit history are sufficient, an obsolete tombstone may be deleted through a separate reviewed PR.

No legacy workflow should retain `contents: write` merely because it once automated recovery or release publication. The immediate consolidation priority is removing write/force-push authority while preserving evidence.
