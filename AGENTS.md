# Project-wide agent operating contract

This file applies to every automated agent and contributor working in this repository.

## One canonical development line

During the current consolidation phase, `integration/canonical` is the single integrated development line and `main` is the stable baseline. Before doing any work, verify the live GitHub refs and read the consolidation status/checklist. Do not start another broad `develop`, `integration`, recovery, or security line in parallel.

After consolidation, all development continues from the canonical integrated state. Short-lived topic branches are allowed when useful for review, but their work must return to the canonical line through validation/review. Never push directly or force-push to `main`.

## Agent autonomy

Agents have broad autonomy to inspect, design, implement, refactor, test, document, harden, research, and coordinate work needed to advance the project. They may decompose work into specialist sub-agents or parallel workstreams when the execution environment supports that capability. Any sub-agent inherits this entire contract and must report its changes/evidence back to the canonical project state.

Autonomy does not override platform permissions, safety boundaries, repository protections, review gates, or explicit owner restrictions. Agents must not merge to `main`, publish releases/tags/packages, rotate credentials, weaken protections, fabricate recovery bytes/evidence, or discard unclassified work unless separately and explicitly authorized.

## Coordination rule

Before creating a branch or PR, check whether an active canonical branch/PR already covers the objective. Prefer updating/continuing existing canonical work over creating overlapping branches. Parallel specialist branches must be narrowly scoped, short-lived, and reconciled into the canonical line; they are never competing sources of truth.

Agents should leave durable handoff information in the repository: objective, affected files, validation performed, PASS/FAIL/BLOCKED/NOT RUN evidence, blockers, and next safe action. Future agents must consume that handoff before continuing.

## Development priorities

1. Finish repository consolidation without losing unique work.
2. Neutralize legacy direct/force-write publication paths and minimize workflow permissions.
3. Establish a complete, provenance-backed, buildable canonical application source.
4. Validate dependencies, supply chain, secrets exposure paths, privacy/network behavior, update mechanisms, packaging/install/uninstall, and recovery behavior.
5. Continue product/UI/customization development from the canonical source only.
6. Keep installation simple and public metadata/screenshots/version claims aligned with verified builds.
7. Test aggressively and fix discovered defects before release preparation.

## Evidence and safety

Use PASS only for checks actually executed successfully against the stated commit/input. Use FAIL for executed failures, BLOCKED for missing/invalid authoritative prerequisites, and NOT RUN for checks not executed. Never convert historical or assumed results into current PASS evidence.

Recovery and validation workflows observe/reconstruct; they do not publish. External or mutable inputs require provenance/integrity controls. Missing source bytes or objects remain blockers rather than being guessed.

## Completion behavior

Agents should keep advancing independently within these boundaries instead of waiting for routine implementation decisions. Escalate only when an action requires owner authorization, credentials, destructive cleanup, release publication, a consequential product decision with no safe default, or when authoritative source/evidence is unavailable.
