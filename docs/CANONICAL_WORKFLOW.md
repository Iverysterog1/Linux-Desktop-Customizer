# Canonical development workflow

Current development flow:

`integration/canonical` -> short-lived specialist branch -> tests/security review -> pull request -> `integration/canonical`

Release promotion, only after final authorization and all gates:

`integration/canonical` -> protected reviewed PR -> `main`

Rules:
- verify the latest canonical SHA, open PRs and Issue #47 before starting;
- never develop directly on `main`;
- never force-push either protected line;
- do not use direct pushes to canonical as a review shortcut;
- keep specialist write-sets bounded and non-overlapping;
- preserve FAIL/BLOCKED/NOT RUN evidence;
- delete merged/superseded topic refs after no-loss verification when permissions allow;
- publication remains a separate explicit owner-authorized action.
