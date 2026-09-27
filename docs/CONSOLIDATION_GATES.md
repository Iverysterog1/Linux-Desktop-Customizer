# Consolidation gates

**G1 Inventory:** every branch/PR classified.

**G2 Security:** every legacy write-capable workflow has a canonical least-privilege disposition.

**G3 Recovery:** incomplete/corrupt source is not promoted; recovery evidence remains reproducible.

**G4 Validation:** canonical diff and CI are reviewed with PASS/FAIL/BLOCKED/NOT RUN stated accurately.

**G5 Cleanup:** only after G1-G4, superseded PRs/branches may be cleaned in a separate explicit operation.

**G6 Development:** feature work restarts from the canonical line after consolidation, not from historical branches.
