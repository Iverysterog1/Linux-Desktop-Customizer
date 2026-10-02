# KDE-008 monitor-aware restore architecture

Status: architecture only. No monitor mutation is implemented by this document.

## Purpose

Monitor geometry and membership can change between preview and apply. A saved layout must not be replayed against a different live topology.

KWin 6 documents a read-only workspace screen list, output geometry and scale information, and a screensChanged signal when the output list changes.

Official references:
- https://develop.kde.org/docs/plasma/kwin/api/
- https://develop.kde.org/docs/plasma/kwin/

## Fail-closed design

A future monitor-aware transaction must keep four concepts separate:

1. normalized live topology;
2. deterministic topology fingerprint;
3. bounded state payload;
4. transaction snapshot and rollback metadata.

Rules:

- Normalize output order before fingerprinting.
- Never use screen index alone as identity.
- Geometry alone is not sufficient identity.
- Missing or ambiguous identity is BLOCKED.
- Added, removed, renamed, or materially changed outputs invalidate preview.
- Geometry or scale changes that affect the payload invalidate preview.
- Observe topology again immediately before mutation.
- The pre-apply fingerprint must equal the fingerprint approved at preview.
- A screensChanged event between preview and apply invalidates approval.
- Default to exact topology compatibility. Any relaxed matching rule needs separate tests and Supervisor approval.

## Required transaction sequence

detect -> observe topology -> preview -> snapshot -> re-observe topology -> compare fingerprint -> apply bounded state -> validate -> rollback/unapply

A stale preview must never be silently recomputed and applied.

Before mutation, all affected state must be journaled. After mutation, validate the intended state and retain the prior snapshot until validation succeeds. On failure, attempt rollback and validate rollback. Rollback failure remains FAIL.

## Evidence status

- PASS: a non-mutating architecture/test component passed, or a future real restore completed with validated evidence.
- FAIL: attempted apply or rollback did not produce its required validated state.
- BLOCKED: topology is missing, ambiguous, changed, or required runtime capability is unavailable.
- NOT RUN: real monitor-specific mutation was not executed.

Architecture PASS must never be presented as real restore PASS.

## Safety boundary

- Keep evidence local by default.
- Collect only topology facts required for safe matching.
- Do not collect screenshots, window contents, user files, or unrelated environment data.
- No downloaded KWin scripts or arbitrary profile commands.
- No privilege elevation or package installation in the restore path.
- New KWin/display command surfaces require explicit allowlisting and Security A review.
- CI, dependency, or packaging changes require Security B review.
- Existing transaction, path-boundary, and rollback guarantees remain mandatory.

## Tests required before implementation is enabled

Automated tests must cover:

- stable fingerprint across output enumeration order;
- exact topology match;
- monitor added or removed;
- ambiguous displays;
- missing identity data;
- geometry or scale change;
- topology change between preview and apply;
- apply failure followed by rollback;
- rollback failure remains FAIL;
- diagnostics expose only bounded topology evidence.

Real Plasma evidence must later cover representative single- and multi-monitor sessions. Wayland and X11 evidence are tracked separately.

## Current decision

KDE-008 architecture: PASS.

Monitor-specific implementation: NOT RUN.

Real Plasma monitor restore: NOT RUN.

No monitor-specific restore code should be enabled until the topology model, matching tests, transaction coverage, and Security A/B review are integrated through the protected-branch process.
