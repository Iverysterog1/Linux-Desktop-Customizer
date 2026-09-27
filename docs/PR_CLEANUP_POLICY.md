# PR and branch cleanup policy

Cleanup is the last consolidation phase, not the first.

A PR may be closed as superseded only after its unique change is present on the canonical integration line or deliberately rejected with rationale. A topic branch may be deleted only after the same determination and after preserving enough PR/commit history for audit.

Do not rewrite historical branch tips merely to make the branch list look cleaner. Do not force-update the canonical integration branch to absorb divergent work. Prefer explicit commits whose provenance is reviewable.
