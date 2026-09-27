# Consolidation validation

Before the consolidation PR is eligible for review as complete:

- compare it with `main` and inspect every changed file;
- verify all changed/new workflows are read-only unless an explicitly reviewed narrower permission is necessary;
- confirm no changed workflow contains direct/force publication to `main`;
- run the centralized workflow regression guard;
- verify recovery blocker scripts fail closed rather than promoting incomplete source;
- cross-check each open PR against the consolidation map;
- record CI PASS/FAIL/NOT RUN truthfully.

A clean branch list is not itself evidence of successful consolidation.
